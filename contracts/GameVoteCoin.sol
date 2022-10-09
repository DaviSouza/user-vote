pragma solidity >=0.4.22 <0.9.0;

contract GameVoteCoin {
    mapping(address => uint256) public balances;

    uint256 public totalSupply = 300000000 * 10**18;
    string public name = "FirstCoin";

    event Transfer(address indexed from, address indexed to, uint256 value);
    mapping(address => mapping(address => uint256)) public allowance;

    constructor() public {
        balances[msg.sender] = totalSupply;
    }

    function balanceOf(address owner) public view returns (uint256) {
        return balances[owner];
    }

    function transfer(address to, uint256 value) public returns (bool) {
        require(balanceOf(msg.sender) >= value, "balance too low");
        balances[to] += value;
        balances[msg.sender] -= value;
        emit Transfer(msg.sender, to, value);
        return true;
    }

    function transferFrom(
        address from,
        address to,
        uint256 value
    ) public returns (bool) {
        require(balanceOf(from) >= value, "balance too low");
        require(allowance[from][msg.sender] >= value, "allowance too low");
        balances[to] += value;
        balances[from] -= value;
        emit Transfer(from, to, value);
        return true;
    }
}
