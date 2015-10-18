/*Package prism is a package that helps you write traffic splitting tools.

An HTTP splitter sits in line on an http request/response path and replays some
portion of traffic to additional backends.

A splitter is defined by one upstream "source" and will repeat some portion of
traffic to a number of "sinks".

Example configuration:

  {
    "Splitters": [
      {
        "Label": "main",
        "Source": "http://localhost:8000",
        "ListenAddr": "localhost:7000",
        "SinkRequestTimeout": 30,
        "WaitForResponse": true,
        "Sinks": [
          {
            "Name": "version a",
            "Addr": "localhost:8001"
          },
          {
            "Name": "version b",
            "Addr": "localhost:8002"
          }
        ]
      }
    ],
    "MonitorAddr": ":7070"
  }
*/
package prism
