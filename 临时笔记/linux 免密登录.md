# linux 免密登录

## 步骤

1. 将 Client 端公钥的内容复制服务器
2. chmod 700 ~/.ssh/
3. cat xxx_rsa.pub >> authorized_keys
4. chmod 600 authorized_keys
5. systemctl restart sshd.service