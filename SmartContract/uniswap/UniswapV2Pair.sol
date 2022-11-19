pragma solidity >=0.5.0;

contract UniswapV2Pair {
    uint public totalSupply;
    mapping(address => uint) public balanceOf;
    address public factory;
    address public token0;
    address public token1;
}