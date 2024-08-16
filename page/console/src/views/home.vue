<template>
    <div class="app-content no-select">
        <div class="head-bar flex mb-10">
            <div class="card tips mr-10">
                <div class="float-right pa-10">{{ tips }}好呀 {{ user.nickname }} 😊</div>
                <div class="text-gray text-small pa-10 border-bottom">
                    <n-time :time="now" />
                </div>
                <div class="pa-10">
                    <div>欢迎回来, 今天要不写点什么?</div>
                </div>
            </div>
            <div class="card readme full-width">
                <div class="flex readme-info pa-10 border-bottom">
                    <n-avatar class="readme-avatar mr-10" size="large" :src="user.avatar" />
                    <div>
                        <div class="flex align-center">
                            <div class="mr-5">{{ user.nickname }}</div>
                            <div class="text-gray text-small">@{{ user.username }}</div>
                        </div>
                        <div class="text-gray text-small lint1">{{ user.bio }}</div>
                    </div>
                </div>
                <n-button strong secondary type="primary" class="full-width" @click="jump('/add')">撰写新文章</n-button>
            </div>
        </div>
        <div class="card">
            <n-button class="float-right submit-btn" :loading="submit" type="primary" @click="addNote">保存</n-button>
            <div class="pa-10">随手记</div>
            <div class="">
                <n-input v-model:value="note" :disabled="submit" placeholder="随手写点什么吧, 别管格式了" type="textarea"
                    size="small" :autosize="{
                        minRows: 20,
                        maxRows: 20,
                    }" />
            </div>
        </div>
    </div>
</template>

<script>
import { Github, Linkedin, Discord } from '@vicons/fa'
import { QqOutlined } from '@vicons/antd'
import { article } from '../plugins/api'

export default {
    name: "Home",
    components: { Github, Linkedin, QqOutlined, Discord },
    data: () => ({
        now: 0,
        tips: '',
        note: '',
        submit: false,
        user: {}
    }),
    methods: {
        init() {
            let user = localStorage.getItem('cache:user');
            if (user != undefined) {
                this.user = JSON.parse(user)
            }
            this.initTimePeriod()
        },
        jump(path) {
            this.$router.push(path)
        },
        addNote() {
            if (this.note.trim().length == 0) {
                window.$message.warning('没有可以保存的内容')
                return false
            }
            window.$dialog.warning({
                title: "操作确认",
                content: "随手记保存后如需修改, 需前往文章列表中查找本篇随手记后进入详情修改, 确认要保存吗?",
                positiveText: "确认保存",
                negativeText: "取消",
                onPositiveClick: () => {
                    this.submit = true
                    article.add(false, false, 1, '随手记 #' + new Date().getTime(), "", 0, null, this.note, 1, null).then(res => {
                        if (res.state) {
                            this.note = ''
                            window.$message.success('保存成功')
                        } else
                            window.$message.warning(res.message ? res.message : '保存失败')
                        this.submit = false
                    }).catch(err => {
                        console.log(err)
                        this.submit = false
                        window.$message.error('保存出错')
                    })
                },
            });
        },
        initTimePeriod() {
            const currentHour = new Date().getHours();
            if (currentHour >= 5 && currentHour < 11) 
                this.tips = '早上';
             else if (currentHour >= 11 && currentHour < 13) 
                this.tips = '中午';
             else if (currentHour >= 13 && currentHour < 18) 
                this.tips = '下午';
             else 
                this.tips = '晚上'
        }
    },
    mounted() {
        this.now = new Date().getTime()
        this.init()
    },
};
</script>

<style scoped>
.tips {
    width: 100%;
}

.readme {
    min-width: 200px;
    max-width: 300px;
    height: 116px;
}

.readme:deep(.social-btn) {
    width: 25%;
}

.submit-btn {
    margin: 4px 4px 0 0;
}

@media only screen and (max-width: 767px) {
    .head-bar {
        display: block;
    }

    .tips {
        margin-bottom: 10px;
    }
}
</style>
