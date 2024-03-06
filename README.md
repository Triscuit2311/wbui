# WIP: Web Based Universal Interface
## An interface for (any, or multiple) native applicatons
## Written in Go with HTMX
I identified that when I am protoyping native applications, sometimes I want runtime controls but do not need a GUI for the finished application or simply don't feel the time investment is justified.
The idea is here is a UI Daemon that serves a basic UI locally. The goal is top provide a universal API for native applications (using the filesystem or a socket) so they can add controls or display data.
