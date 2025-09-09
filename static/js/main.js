document.getElementById('signInButton').addEventListener('click', ClickButton);

function CheckBoxInner() {
    const username = document.getElementById('username').value;
    const email = document.getElementById('email').value;

    if (!username || !email) {
        return false; // 表单未填写完整
    }
    return true; // 表单已填写完整
}

function ClickButton() {
    if (CheckBoxInner()) {
        const signInButton = document.getElementById('signInButton');
        signInButton.classList.add('button-success'); // 添加成功样式类
        signInButton.parentElement.classList.remove('button-section'); // 移除初始按钮样式类
        signInButton.textContent = '签到成功';
        signInButton.disabled = true; // 禁用按钮，防止重复点击
    }
}