const loginBtn = document.getElementById('login-btn');
const loginInput = document.getElementById('login');
const passwordInput = document.getElementById('password');
const errorMsg = document.getElementById('login-error-msg');

loginBtn.addEventListener('click', async (e) => {
    e.preventDefault();
    const login = loginInput.value;
    const password = passwordInput.value;
    if (login === '' || password === '') {
        errorMsg.style.opacity = 1;
    }

    let authURL = 'http://localhost:7000/auth'

    // get query params from cookie to pass it if login succeeds
    const params = getCookie("params")
    if (params !== "") {
        document.cookie = `${params}`
    }

    try {
        const response = await loginUser(login, password);
        const jwt = response['jwt']

        // Set JWT token as cookie
        document.cookie = `jwt=Bearer ${jwt}; path=/`;
        window.location.href = authURL
    } catch (error) {
        console.error(error)
    }
});

async function loginUser(login, password) {
    const response = await fetch("http://localhost:7000/authentication/auth", {
        method: "POST",
        headers: {
            "Content-Type": "application/json;charset=utf-8"
        },
        body: JSON.stringify({login, password}),
        credentials: "same-origin",
    })

    if (!response.ok) {
        throw new Error("login failed")
    }

    return await response.json();
}

function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length > 1) return parts[1].split(';').shift();
}
