"use strict";
$('#login').click(() => {
    let password = $('#password').val();
    login(password);
});
const login = async (password) => {
    try {
        const res = await fetch("/login", {
            method: 'POST',
            body: JSON.stringify(password),
            headers: { "content-type": "application/json" }
        });
        if (res.ok) {
            setTimeout(window.location.reload.bind(window.location), 1000);
        }
    }
    catch (error) {
        console.log(error);
    }
};
