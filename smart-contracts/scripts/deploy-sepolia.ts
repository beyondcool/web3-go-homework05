import { network } from "hardhat";

const { ethers } = await network.create({
  network: "sepolia",
});

console.log("Deploying Counter to Sepolia...");

const counter = await ethers.deployContract("Counter");
await counter.waitForDeployment();

const address = await counter.getAddress();
console.log("Counter deployed to:", address);
console.log("Transaction hash:", counter.deploymentTransaction()!.hash);