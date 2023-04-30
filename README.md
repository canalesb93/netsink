# Netsink

Netsink is a network simulation tool that allows you to configure your network connection to mimic slow and unreliable connections. It is particularly useful for testing applications under various network conditions to ensure they function correctly.

## Installation

To install Netsink, follow these steps:

1. Ensure you have Go installed on your machine. If not, follow the instructions at https://golang.org/doc/install to set up Go.

2. Clone the Netsink repository to your local machine:

   ```
   git clone https://github.com/username/netsink.git
   ```

   Replace `username` with the correct username or organization name.

3. Navigate to the Netsink directory:

   ```
   cd netsink
   ```

4. Install the Netsink binary:

   ```
   go install ./cmd/netsink
   ```

   Make sure the `$GOPATH/bin` directory is in your system's `PATH` environment variable.

After the installation, you should be able to run the `netsink` command from anywhere on your machine.

## Usage

Run `netsink` to interactively configure your network connection with parameters such as delay and packet loss percentage. To reset the network conditions to normal, run `netsink --reset` or `netsink -r`.

For more detailed usage information, run `netsink -h` or `netsink --help`.

## Future Enhancements

- Add support for simulating network latency and packet loss on a per-IP or per-subnet basis.
- Implement an interactive mode with a terminal user interface (TUI) for a better user experience.
- Add the ability to import and export network profiles for easy sharing and reusability.
- Integrate with network monitoring tools to provide real-time feedback on simulated network conditions.
- Add support for simulating network conditions based on preset templates (e.g., 3G, 4G, or satellite connections).