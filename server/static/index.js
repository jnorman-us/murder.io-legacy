console.log('test');

if (!WebAssembly.instantiateStreaming) { // polyfill
    WebAssembly.instantiateStreaming = async (resp, importObject) => {
        const source = await (await resp).arrayBuffer();
        return await WebAssembly.instantiate(source, importObject);
    };
}

const go = new Go();
WebAssembly.instantiateStreaming(fetch('client.wasm'), go.importObject).then(res => {
    go.run(res.instance)
});

setTimeout(function() {
    var username = prompt("What is your username?", "");
    connectToServer(username);
}, 1000);