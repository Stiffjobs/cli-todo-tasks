# Todo CLI Tool

This CLI tool is a simple task management application built using Go and the Cobra library. It allows you to manage your todo list directly from the command line.

## Features

- **List Tasks**: View your incomplete tasks.
- **Add Tasks**: Add new tasks to your list.
- **Delete Tasks**: Delete tasks by their ID.
- **Show Completed Tasks**: View all completed tasks with an additional flag.

## Usage

### List Incomplete Tasks

To list all your incomplete tasks, use the following command:

```bash
tasks list
```

### List All Tasks (Including Completed)

To list all tasks, including completed ones, add the `-a` flag:

```bash
tasks list -a
```

### Add a New Task

To add a new task, use the following command:

```bash
task add "Your new task here"
```

### Delete a Task by ID

To delete a task, you can use the task's ID as follows:

```bash
tasks delete <ID>
```

Replace `<ID>` with the actual ID of the task you want to delete.

## Examples

- **Adding a Task**:

  ```bash
  tasks add "Buy groceries"
  ```

- **Listing All Tasks**:

  ```bash
  tasks list -a
  ```

- **Deleting a Task**:

  ```bash
  tasks delete 2
  ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
