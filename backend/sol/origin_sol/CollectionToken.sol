// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract CollectionToken is IERC20 {

    using SafeMath for uint256;

    mapping (address => uint256) private _balances;

    mapping (address => mapping(address => uint256)) private _allowed;

    uint256 private _totalSupply;

    string public symbol;

    address private _owner;

    constructor(string memory _sym) public {
        symbol = _sym;
        _owner = msg.sender;
    }

    function totalSupply() override public view returns (uint256) {
        return _totalSupply;
    }

    function balanceOf(address owner) override public view returns (uint256) {
        return _balances[owner];
    }

    function approve(address spender, uint256 value) override public returns (bool) {

        //不能往零地址发
        require(spender != address(0));

        //授权的值不能超过拥有的值
        require(value <= _balances[msg.sender]);


        //更改之
        _allowed[msg.sender][spender] = value;
        emit Approval(msg.sender, spender ,value);
        return true;
    }

    //某账户被允许使用的代币
    function allowance(address owner, address spender) override public view returns (uint256) {
        return _allowed[owner][spender];
    }

    //从from账户转账给to账户
    function transferFrom(address from, address to, uint256 value) override public returns (bool) {

        //用户余额充足
        require(value <= _balances[from]);
        //用户授权余额充足
        require(value <= _allowed[from][msg.sender]);
        require(to != address(0));

        _balances[from] = _balances[from].sub(value);
        _balances[to] = _balances[to].add(value);
        _allowed[from][msg.sender] = _allowed[from][msg.sender].sub(value);
        emit Transfer(from,to,value);
        return true;
    }

    //从自己的账户转账给to账户
    function transfer(address to, uint256 value) override public returns (bool) {
        //用户余额充足
        require(value <= _balances[msg.sender]);
        require(to != address(0));

        _balances[msg.sender] = _balances[msg.sender].sub(value);
        _balances[to] = _balances[to].add(value);
        emit Transfer(msg.sender,to,value);
        return true;
    }

    //获取代币
    function mint(address to, uint256 value) public {
        require(msg.sender == _owner);
        _totalSupply = _totalSupply.add(value);
        _balances[to] = _balances[to].add(value);
        emit Transfer(address(0), to, value);
    }

}

