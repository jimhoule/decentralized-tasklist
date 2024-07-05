// SPDX-License-Identifier: GPL-3.0

pragma solidity >=0.8.2 <0.9.0;

contract TaskList {
    struct Info {
        string name;
        string description;
        string email;
    }

    struct Task {
        string id;
        string name;
        string description;
        bool isDone;
    }

    struct PartialTask {
        string id;
        string name;
        bool isDone;
    }

    address private owner;
    string private name;
    string private description;
    string private email;
    Task[] private tasks;

    constructor(string memory _name, string memory _description, string memory _email) {
        // NOTE: Sets account deploying the contract as owner (msg.sender returns the address of the account trying to interact with the contract)
        owner = msg.sender;
        name = _name;
        description = _description;
        email = _email;
    }

    // NOTE: Restricts access to contract owner only
    modifier isOwner() {
        require(owner == msg.sender);
        _;
    }

    function getInfo() public isOwner view returns (Info memory) {
        return Info(name, description, email);
    }

    function getTask(string memory _id) public isOwner view returns (Task memory) {
        Task memory task;

        for (uint i = 0; i < tasks.length; i++) {
            bool isTaskFound = compareStrings(tasks[i].id, _id);

            if (isTaskFound) {
                task = tasks[i];
            }
        }

        return task;
    }

    function getTasks() public isOwner() view returns (PartialTask[] memory) {
        PartialTask[] memory partialTasks = new PartialTask[](tasks.length);

        for (uint i = 0; i < tasks.length; i++) {
            Task memory task = tasks[i];
            partialTasks[i] = PartialTask(task.id, task.name, task.isDone);
        }

        return partialTasks;
    }

    function addTask(
        string memory _id,
        string memory _name,
        string memory _description
    ) public isOwner {
        tasks.push(Task(_id, _name, _description, false));
    }

    function updateTask(
        string memory _id,
        string memory _name,
        string memory _description,
        bool _isDone
    ) public isOwner {
        uint indexToUpdate = getTaskIndex(_id);

        tasks[indexToUpdate].name = _name;
        tasks[indexToUpdate].description = _description;
        tasks[indexToUpdate].isDone = _isDone;
    }

    function deleteTask(string memory _id) public isOwner {
        uint indexToRemove = getTaskIndex(_id);

        for (uint i = indexToRemove; i < tasks.length - 1; i++) {
            tasks[i] = tasks[i + 1];
        }

        tasks.pop();
    }

    function getTaskIndex(string memory _id) private view returns (uint) {
        uint index = 0;

        for (uint i = 0; i < tasks.length; i++) {
            bool isTaskFound = compareStrings(tasks[i].id, _id);

            if (isTaskFound) {
                index = i;
                break;
            }
        }

        return index;
    }

    function compareStrings(string memory str1, string memory str2) private pure returns (bool) {
        bytes memory b1 = bytes(str1);
        bytes memory b2 = bytes(str2);

        if (b1.length != b2.length) {
            return false;
        }

        for (uint256 i = 0; i < b1.length; i++) {
            if (b1[i] != b2[i]) {
                return false;
            }
        }

        return true;
    }
}