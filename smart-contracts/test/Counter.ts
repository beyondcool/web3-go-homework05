import { expect } from "chai";
import { network } from "hardhat";

const { ethers } = await network.create();

describe("Counter", function () {
  it("Should emit the Increment event when calling the inc() function", async function () {
    const counter = await ethers.deployContract("Counter");

    await expect(counter.inc()).to.emit(counter, "IncrementEvent").withArgs(1n);
  });

  it("The sum of the Increment events should match the current value", async function () {
    const counter = await ethers.deployContract("Counter");
    const deploymentBlockNumber = await ethers.provider.getBlockNumber();

    // run a series of increments
    for (let i = 1; i <= 10; i++) {
      await counter.incBy(i);
    }

    const events = await counter.queryFilter(
      counter.filters.IncrementEvent(),
      deploymentBlockNumber,
      "latest",
    );

    // check that the aggregated events match the current value
    let total = 0n;
    for (const event of events) {
      total += event.args.by;
    }

    expect(await counter.num()).to.equal(total);
  });

  describe("donate()", function () {
    it("Should accept ETH and emit DonateEvent", async function () {
      const [_, user1] = await ethers.getSigners();
      const counter = await ethers.deployContract("Counter");
      const deployBlock = await ethers.provider.getBlockNumber();

      const amount = ethers.parseEther("1.0");

      // donate
      const tx = await counter.connect(user1).donate({ value: amount });
      await expect(tx)
        .to.emit(counter, "DonateEvent")
        .withArgs(user1.address, amount);
    });

    it("Should revert when donating 0 ETH", async function () {
      const [_, user1] = await ethers.getSigners();
      const counter = await ethers.deployContract("Counter");

      await expect(
        counter.connect(user1).donate({ value: 0n })
      ).to.be.revertedWith("donate: amount should be positive");
    });
  });

  describe("withdraw()", function () {
    it("Should allow owner to withdraw and emit WithdrawEvent", async function () {
      const [owner, user1] = await ethers.getSigners();
      const counter = await ethers.deployContract("Counter");

      const amount = ethers.parseEther("1.0");
      await counter.connect(user1).donate({ value: amount });

      const tx = await counter.connect(owner).withdraw();
      await expect(tx)
        .to.emit(counter, "WithdrawEvent")
        .withArgs(owner.address, amount);

      // contract balance should be 0 after withdrawal
      expect(await ethers.provider.getBalance(await counter.getAddress())).to.equal(0n);
    });

    it("Should revert when non-owner tries to withdraw", async function () {
      const [_, user1] = await ethers.getSigners();
      const counter = await ethers.deployContract("Counter");

      await expect(
        counter.connect(user1).withdraw()
      ).to.be.revertedWithCustomError(counter, "OwnableUnauthorizedAccount")
        .withArgs(user1.address);
    });

    it("Should revert when totalBalance is 0", async function () {
      const [owner] = await ethers.getSigners();
      const counter = await ethers.deployContract("Counter");

      await expect(
        counter.connect(owner).withdraw()
      ).to.be.revertedWith("withdraw: totalBalance should be positive");
    });
  });
});