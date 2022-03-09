function createScene(engineUpdate, drawUpdate, centerUpdate) {
    const width = window.innerWidth;
    const height = window.innerHeight;

    const stats = new Stats();
    stats.showPanel(0);
    const scene = new THREE.Scene();
    const camera = new THREE.OrthographicCamera(0, 0, 0, 0, -1000, 1000);
    const aspect = width / height;
    const d = 200;
    camera.left = aspect * -d;
    camera.right = aspect * d;
    camera.top = 1 * d;
    camera.bottom = 1 * -d;
    camera.updateProjectionMatrix();

    const sun_light = new THREE.DirectionalLight(0xffffff, 3);
    const sun_light_target = new THREE.Object3D();
    sun_light.target = sun_light_target;
    sun_light.castShadow = true;
    sun_light.shadow.mapSize.copy(new THREE.Vector2(2000, 2000));
    sun_light.shadow.camera.zoom = .01;
    sun_light.shadow.camera.near = -2000;
    sun_light.shadow.camera.far = 2000;
    scene.add(sun_light);
    scene.add(sun_light_target);

    const ambient_light = new THREE.AmbientLight(0xFFF6DA, 1);
    scene.add(ambient_light);

    const planeWidth = 1000;
    const planeHeight = 1000;
    const ground = new THREE.Mesh(new THREE.PlaneBufferGeometry(planeWidth, planeHeight, 32, 32), new THREE.MeshStandardMaterial({
        color: 0xffffff,
    }));
    ground.rotation.set(-Math.PI / 2, 0, 0, "YXZ");
    ground.position.set(250, 0, 250);
    ground.receiveShadow = true;
    ground.castShadow = true;
    scene.add(ground);

    const renderer = new THREE.WebGLRenderer();
    renderer.physicallyCorrectLights = true;
    renderer.shadowMap.enabled = true;
    renderer.setSize(width, height);

    camera.position.set(100, 100, 100);
    camera.lookAt(new THREE.Vector3(0, 0, 0));

    const entities = new Map();

    document.body.appendChild(stats.dom);
    document.body.appendChild(renderer.domElement);

    function animate() {
        stats.begin();
        if(go.exited) return;
        requestAnimationFrame(animate);
        engineUpdate(); // golang function
        parseEngineDraw();
        renderer.render(scene, camera);
        stats.end();
    }
    animate();

    function parseEngineDraw() {
        const center = JSON.parse(centerUpdate());
        camera.position.x = center.Position.X + 100;
        camera.position.z = center.Position.Y + 100;

        const centerVec3 = new THREE.Vector3(center.Position.X, 0, center.Position.Y);
        sun_light.position.set(-100, 100, -200);
        sun_light.position.add(centerVec3);
        sun_light_target.position.set(0, 0, 0);
        sun_light_target.position.add(centerVec3);

        const objects = JSON.parse(drawUpdate()); // golang function
        const missingIDs = new Set(entities.keys());
        for(const [id, object] of Object.entries(objects)) {
            missingIDs.delete(id);
            if(entities.has(id)) {
                const mesh = entities.get(id);
                mesh.position.x = object.Position.X;
                mesh.position.z = object.Position.Y;

                mesh.rotation.y = -object.Angle;
            } else {
                const geometry = new THREE.BoxGeometry(10, 10, 10);
                const color = object.Color;
                const material = new THREE.MeshStandardMaterial({
                    color: new THREE.Color(color.R, color.G, color.B),
                });
                const mesh = new THREE.Mesh(geometry, material);
                mesh.position.x = object.Position.X;
                mesh.position.z = object.Position.Y;
                mesh.position.y = 5;

                mesh.rotation.y = object.Angle;

                mesh.receiveShadow = true;
                mesh.castShadow = true;

                scene.add(mesh);
                entities.set(id, mesh);
            }
        }

        for(const id of missingIDs) {
            const mesh = entities.get(id);
            scene.remove(mesh);
            entities.delete(id);
        }
    }
}