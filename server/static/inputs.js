function initInputs(setInputs) {
    const inputs = {
        forward: false,
        backward: false,
        left: false,
        right: false,
        attackClick: false,
        rangedClick: false,
        special: false,
        direction: 0,
    }
    const keyToAction = {
        '0': 'attackClick',
        '2': 'rangedClick',
        'w': 'forward',
        'a': 'left',
        's': 'backward',
        'd': 'right',
        ' ': 'special',
    }

    window.addEventListener('mousedown', function(e) {
        handleClickOrPress(e.button, true);
    });
    window.addEventListener('mouseup', function(e) {
        handleClickOrPress(e.button, false);
    });
    window.addEventListener('keydown', function(e) {
        handleClickOrPress(e.key, true);
    });
    window.addEventListener('keyup', function(e) {
        handleClickOrPress(e.key, false);
    });
    window.addEventListener('contextmenu', function(e) {
        e.preventDefault();
    });

    function handleClickOrPress(key, active) {
        const action = keyToAction[key];
        if(action != null && !go.exited) {
            inputs[action] = active;
            setInputs(inputs);
        }
    }
}