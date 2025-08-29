const registerForm = document.getElementById('registerForm');
const loginForm = document.getElementById('loginForm');

registerForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    const res = await fetch('/api/register', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            username: document.getElementById('regUsername').value,
            password: document.getElementById('regPassword').value
        })
    });
    alert(await res.text());
});

loginForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    const res = await fetch('/api/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            username: document.getElementById('loginUsername').value,
            password: document.getElementById('loginPassword').value
        })
    });
    alert(await res.text());
});
