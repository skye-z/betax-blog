<template>
    <div class="flex align-center justify-center full-width text-gray pa-10">
        <div class="mr-5">Copyright &copy; {{ year }}</div>
        <a :href="'https://github.com/' + info.username" rel="noopener" target="_blank">
            <n-button text class="text-gray">{{ info.nickname }}</n-button>
        </a>
        <div class="ml-10">&hearts;</div>
        <div class="ml-10 mr-5">
            Built on
            <a href="https://github.com/skye-z/betax-blog" rel="noopener" target="_blank">
                <n-button text class="text-gray">BetaX Blog</n-button>
            </a>
        </div>
        <div class="update" v-if="version.tag_name && info.version && version.tag_name != info.version" @click="update">
            <n-icon>
                <ArrowCircleUp12Filled />
            </n-icon>
        </div>
    </div>
</template>
<script>
import { ArrowCircleUp12Filled } from '@vicons/fluent'
import { setting } from '../plugins/api'

export default {
    name: "FootBar",
    components: { ArrowCircleUp12Filled },
    props: {
        info: {
            type: Object,
            default: () => { }
        }
    },
    data: () => ({
        year: '',
        version: {}
    }),
    methods: {
        check() {
            setting.checkVersion().then(res => {
                if (res.state) this.version = res.data
            }).catch(err => {
                console.log(err)
            })
        },
        update() {
            window.$dialog.info({
                title: '有新版本可用',
                content: `当前版本: ${this.info.version}\t最新版本: ${this.version.tag_name}\n版本名称: ${this.version.name}\n更新说明:\n${this.version.body}`,
                positiveText: '开始更新',
                negativeText: '取消',
                onPositiveClick: () => {
                    let tips = window.$message.loading('正在更新中', { duration: 1000 * 180 });
                    setting.updateVersion().then(res => {
                        if (res.state) {
                            tips.type = "success";
                            tips.content = "更新成功, 请手动刷新页面";
                        } else {
                            tips.type = "warning";
                            tips.content = res.message ? res.message : "更新失败";
                        }
                        setTimeout(() => {
                            tips.destroy()
                        }, 2000);
                    }).catch(err => {
                        console.log(err)
                        tips.type = "warning";
                        tips.content = "更新出错";
                        setTimeout(() => {
                            tips.destroy()
                        }, 2000);
                    })
                }
            })
        }
    },
    mounted() {
        this.year = new Date().getFullYear();
        this.check()
    },
};
</script>
<style scoped>
.update {
    line-height: 16px;
    font-size: 16px;
    cursor: pointer;
    color: red;
}
</style>