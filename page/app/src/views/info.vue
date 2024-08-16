<template>
    <div class="app-content no-select">
        <div class="card">
            <div v-if="loading" class="loading-block"><n-spin /></div>
            <div class="title pa-10 border-bottom">
                <div>{{ info.title }}</div>
                <div class="flex text-small text-gray">
                    <div class="mr-5">发布于</div>
                    <n-time :time="info.releaseTime" />
                    <template v-if="info.releaseTime">
                        <div class="mr-5">, 上次更新是</div>
                        <n-time :time="info.releaseTime" type="relative" />
                    </template>
                </div>
            </div>
            <div class="flex align-center pa-10 border-bottom">
                <n-tag class="mr-5" size="small" :bordered="false" type="success">
                    {{ classMapping[info.class] }}
                </n-tag>
                <template v-if="info.tags">
                    <n-tag class="mr-5" size="small" v-for="tag in info.tags" :bordered="false" type="info">
                        {{ tag.name }}
                    </n-tag>
                </template>
                <n-tag class="mr-5" size="small" v-if="info.isBanner" :bordered="false" type="warning">轮播</n-tag>
                <n-tag class="mr-5" size="small" v-if="info.isUp" :bordered="false" type="error">置顶</n-tag>
            </div>
            <div v-if="info.abstract" class="abstract-box pa-10">
                <div class="abstract-start">“</div>
                <div class="abstract-content">{{ info.abstract }}</div>
                <div class="abstract-end">”</div>
            </div>
            <div class="content border-bottom" v-html="info.content"></div>
            <div class="pa-20 text-center">本文版权归本站所有, 未经许可严禁转载</div>
        </div>
    </div>
</template>

<script>
import { article } from '../plugins/api'
import { marked } from 'marked';

export default {
    name: "Info",
    data: () => ({
        id: 0,
        info: {},
        loading: true,
        classMapping: {},
        tags: [],
    }),
    methods: {
        init() {
            if (this.$route.params.id)
                this.id = parseInt(this.$route.params.id);
            this.initClass();
            this.initTags();
            this.initData();
        },
        initClass() {
            this.classMapping = {};
            article.getClass().then(res => {
                if (res.state) {
                    let list = res.data
                    for (let i in list) {
                        this.classMapping[list[i].id] = list[i].name
                    }
                } else {
                    window.$message.warning('获取分类列表失败');
                }
            }).catch(err => {
                window.$message.warning("发生意料之外的错误");
            })
        },
        initTags() {
            this.tags = [];
            article.getTags().then(res => {
                if (res.state) this.tags = res.data
                else {
                    window.$message.warning('获取标签列表失败');
                }
            }).catch(err => {
                window.$message.warning("发生意料之外的错误");
            })
        },
        initData() {
            article.getInfo(this.id).then(res => {
                if (res.state) {
                    document.title = res.data.title + document.title.substring(document.title.lastIndexOf(' -'))
                    res.data.content = marked(res.data.content)
                    this.info = res.data
                }
                else window.$message.warning('获取文章详情失败');
                this.loading = false;
            }).catch(err => {
                console.log(err)
                this.loading = false;
            })
        },
    },
    mounted() {
        this.init()
    },
};
</script>

<style scoped>
.title {
    font-size: 18px;
    font-weight: bold;
}

.item-time {
    width: 33.33%;
}

.abstract-box{
    position: relative;
    background-color: var(--color-light-3);
}

.dark .abstract-box{
    background-color: var(--color-dark-1);
}

.abstract-start {
    font-size: 48px;
    line-height: 0;
    color: #999;
    position: absolute;
    left: 10px;
    top: 25px;
}

.abstract-end {
    font-size: 48px;
    line-height: 0;
    text-align: right;
    color: #999;
}

.abstract-content {
    font-style: italic;
    padding: 10px 20px;
}

.content{
    margin-top: 10px;
    padding: 0 20px;
}

.content img{
    max-width: 100%;
    margin: 0 auto;
}
</style>