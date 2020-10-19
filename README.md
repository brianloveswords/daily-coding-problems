# daily coding problems

Roll the dice, get a language and a problem.

## Prerequisites

- Docker
- VS Code

This uses [Development Containers](https://code.visualstudio.com/docs/remote/containers) to isolate each language.

## Usage

- Run `bin/roll`
  - If you don't have python3 and really don't want to install it:

    ```shell
    docker run --rm -v ${PWD}:/dcp brianloveswords/dcp-roll
    ```

- Do `code <language>`. VS Code will prompt to re-open in a container.
- Write & test the solution.
  - It probably makes sense to create a new directory for the problem and `code <directory>` to get a container workspace where the root is specific to that solution.
  - Start the language container first, then create the directory, then start the solution container.
