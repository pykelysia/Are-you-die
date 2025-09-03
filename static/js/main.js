document.getElementById('signInButton').addEventListener('click', function() {
    const username = document.getElementById('username').value;
    const email = document.getElementById('email').value;

    if (!username || !email) {
        alert('请输入用户名和邮箱！');
        return;
    }

    const imageSection = document.getElementById('imageSection');
    imageSection.innerHTML = '<img src="/static/IMG_3124.png" alt="签到成功">';

    alert('签到成功！');
});
