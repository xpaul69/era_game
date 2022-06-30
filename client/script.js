const socket = new WebSocket("ws://localhost:8080/ws");

socket.onopen = () => {
        console.log("Connected to websocket");
        socket.send("Hello from client");
}
