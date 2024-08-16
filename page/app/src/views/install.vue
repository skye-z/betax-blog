<template>
    <div class="install-box no-select">
        <div class="install-title">初始化配置</div>
        <div class="card pa-10">
            <n-form ref="formRef" :model="config" :rules="rules" label-placement="left" label-width="85px">
                <div class="form-title pb-10">SSL 加密</div>
                <n-form-item label="证书地址" path="sslCert">
                    <n-input v-model:value="config.sslCert" placeholder="建议使用绝对路径" />
                </n-form-item>
                <n-form-item label="公钥地址" path="sslKey">
                    <n-input v-model:value="config.sslKey" placeholder="建议使用绝对路径" />
                </n-form-item>
                <n-button class="form-title pb-10" text @click="openGithub">Github 登录</n-button>
                <n-form-item label="应用编号" path="githubClientId">
                    <n-input v-model:value="config.githubClientId" placeholder="OAuth2 Client Id" />
                </n-form-item>
                <n-form-item label="调用密钥" path="githubClientSecret">
                    <n-input v-model:value="config.githubClientSecret" placeholder="OAuth2 Client Secret" />
                </n-form-item>
                <n-form-item label="回调地址" path="githubRedirectUrl">
                    <n-input v-model:value="config.githubRedirectUrl"
                        placeholder="http://localhost:9800/oauth2/callback" />
                </n-form-item>
                <div class="form-title pb-10">AI 摘要</div>
                <n-form-item label="服务地址" path="aiSserver">
                    <n-input v-model:value="config.aiSserver" placeholder="请使用OpenAI或与其兼容的服务" />
                </n-form-item>
                <n-form-item label="调用密钥" path="aiSecret">
                    <n-input v-model:value="config.aiSecret" placeholder="请使用OpenAI或与其兼容的服务" />
                </n-form-item>
                <n-form-item label="模型名称" path="aiModel">
                    <n-input v-model:value="config.aiModel" placeholder="请使用OpenAI或与其兼容的服务" />
                </n-form-item>
                <n-form-item label="会话设置" path="aiSetting">
                    <n-input v-model:value="config.aiSetting" type="textarea" :autosize="{
                        minRows: 3,
                        maxRows: 5,
                    }" placeholder="请使用OpenAI或与其兼容的服务" />
                </n-form-item>
            </n-form>
            <div class="text-right">
                <n-button type="primary" @click="save">完成配置</n-button>
            </div>
        </div>
    </div>
</template>

<script>
import { common } from '../plugins/api'

export default {
    name: "Install",
    data: () => ({
        config: {
            sslCert: '',
            sslKey: '',
            githubClientId: '',
            githubClientSecret: '',
            githubRedirectUrl: '',
            aiSserver: '',
            aiSecret: '',
            aiModel: '',
            aiSetting: ''
        },
        rules: {
            githubClientId: {
                required: true,
                message: '请输入 Github 应用编号',
                trigger: 'blur'
            },
            githubClientSecret: {
                required: true,
                message: '请输入 Github 调用密钥',
                trigger: 'blur'
            },
            githubRedirectUrl: {
                required: true,
                message: '请输入 Github 回调地址',
                trigger: 'blur'
            },
        }
    }),
    methods: {
        save() {
            this.$refs.formRef.validate(errors => {
                if (errors) return false

                common.install(this.config).then(res => {
                    if (res.state) {
                        window.$message.success('配置成功');
                        this.$router.push('/').then(() => {
                            location.reload();
                        });
                    } else window.$message.warning('配置失败');
                    this.loading = false
                }).catch(err => {
                    window.$message.warning("发生意料之外的错误");
                    this.loading = false
                })
            })
        },
        openGithub() {
            window.open('https://github.com/settings/developers')
        }
    },
    mounted() {
        this.config.githubRedirectUrl = location.origin + '/oauth2/callback'
    },
};
</script>

<style scoped>
.install-box {
    width: 500px;
    margin: 0 auto;
}

.install-title {
    padding: 30px 0 10px;
    font-size: 32px;
    font-weight: bold;
    text-align: center;
}

.form-title {
    font-weight: bold !important;
}
</style>
