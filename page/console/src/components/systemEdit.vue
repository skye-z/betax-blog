<template>
    <n-drawer v-model:show="show" :width="370" placement="right">
        <n-drawer-content title="系统设置">
            <div v-if="loading" class="loading-block"><n-spin /></div>
            <n-form ref="formRef" :model="config" :rules="rules" label-placement="left" label-width="85px">
                <div class="form-title pb-10">SSL 加密</div>
                <n-form-item label="证书地址" path="sslCert">
                    <n-input v-model:value="config.sslCert" placeholder="建议使用绝对路径" />
                </n-form-item>
                <n-form-item label="公钥地址" path="sslKey">
                    <n-input v-model:value="config.sslKey" placeholder="建议使用绝对路径" />
                </n-form-item>
                <div class="form-title pb-10">Github 登录</div>
                <n-form-item label="应用编号" path="githubClientId">
                    <n-input v-model:value="config.githubClientId" placeholder="OAuth2 Client Id" />
                </n-form-item>
                <n-form-item label="调用密钥" path="githubClientSecret">
                    <n-input v-model:value="config.githubClientSecret" placeholder="OAuth2 Client Secret" />
                </n-form-item>
                <n-form-item label="回调地址" path="githubRedirectUrl">
                    <n-input v-model:value="config.githubRedirectUrl" placeholder="OAuth2 Redirect Url" />
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
            <template #footer>
                <n-button type="primary" @click="save">保存</n-button>
            </template>
        </n-drawer-content>
    </n-drawer>
</template>
<script>
import { setting } from '../plugins/api'

export default {
    name: "SystemEdit",
    data: () => ({
        show: false,
        loading: false,
        config: {},
        rules: {}
    }),
    methods: {
        open() {
            this.loading = true
            this.init()
            this.show = true
        },
        init() {
            setting.getConfig().then(res => {
                if (res.state) this.config = res.data
                else window.$message.warning('列表加载失败');
                this.loading = false
            }).catch(err => {
                window.$message.warning("发生意料之外的错误");
                this.loading = false
            })
        },
        save() {
            setting.updateConfig(this.config).then(res => {
                if (res.state) {
                    this.init();
                    window.$message.success('设置保存成功');
                }
                else window.$message.warning('设置保存失败');
            }).catch(err => {
                window.$message.warning("发生意料之外的错误");
            })
        }
    }
};
</script>
<style scoped>
.form-title {
    font-weight: bold;
}
</style>