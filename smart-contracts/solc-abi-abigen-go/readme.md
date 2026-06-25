# abi

cd smart-contract/

smart-contracts (main)$ solc --base-path ./contracts --include-path ./node_modules  --abi ./contracts/Counter.sol

# go

cd solc-abi-abigen-go/


abigen --abi=Counter.abi --pkg=counter --out=Counter.go



