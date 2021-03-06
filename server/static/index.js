golangEngine = {
    connectToServer: () => {},
    engineUpdate: () => {},
    drawUpdate: () => {},
    centerUpdate: () => {},
    setInputs: () => {},

    objects: {},
};

if (!WebAssembly.instantiateStreaming) { // polyfill
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    };
}

const go = new Go();
WebAssembly.instantiateStreaming(fetch('client.wasm'), go.importObject).then(res => {
    go.run(res.instance).then(r => {
        console.log("Go stopped running");
    });
    handleLogin(golangEngine.connectToServer);
    createScene(golangEngine.engineUpdate, golangEngine.drawUpdate, golangEngine.centerUpdate);
    initInputs(golangEngine.setInputs);
});

function handleLogin(connectToServer) {
    const username = prompt("What is your username? (Press enter to play as \"Beta Tester\")", "Beta Tester");
    const location = window.location;
    const hostname = location.hostname;
    const matches = location.href.matchAll(/:([\d]{2,4})/g);

    let port = 80;
    for(const match of matches) {
        port = parseInt(match[1]);
    }
    connectToServer(hostname, port, username); // golang function
}




