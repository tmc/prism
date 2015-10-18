/*Package prism is a package that helps you write traffic splitting tools.

An HTTP splitter sits in line on an http request/response path and replays some
portion of traffic to additional backends.

A splitter is defined by one upstream "source" and will repeat some portion of
traffic to a number of "sinks".

Example configuration:

  {
    "MonitorAddr": ":7070",
    "Splitters": [
      {
        "Sinks": [
          {
            "Addr": "localhost:8001",
            "Name": "version a"
          },
          {
            "Addr": "localhost:8002",
            "Name": "version b"
          }
        ],
        "WaitForResponse": true,
        "Source": "localhost:8000",
        "Addr": "localhost:7000",
        "Label": "main"
      }
    ]
  }
*/
package prism
