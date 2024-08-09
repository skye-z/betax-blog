<template>
    <div class="app-content no-select">
        <div class="loading" v-if="loading">
            <n-spin />
        </div>
        <div class="card mb-10 pa-10">
            关键词筛选、状态筛选
        </div>
        <div class="card mb-10 article-list">
            <div v-for="(item, index) in list" class="article-item pa-10" @click="toEdit(item.id)" :class="{ 'border-top': index != 0 }">
                <n-button-group class="float-right">
                    <n-button quaternary circle>
                        <template #icon>
                            <n-icon>
                                <ArrowAutofitUp24Filled />
                            </n-icon>
                        </template>
                    </n-button>
                    <n-button quaternary circle>
                        <template #icon>
                            <n-icon>
                                <CameraSwitch24Filled />
                            </n-icon>
                        </template>
                    </n-button>
                    <n-button quaternary circle>
                        <template #icon>
                            <n-icon>
                                <Delete16Regular />
                            </n-icon>
                        </template>
                    </n-button>
                </n-button-group>
                <div class="mb-5 title">{{ item.title }}</div>
                <div>
                    <n-tag class="mr-5" size="small" :bordered="false" type="success">
                        {{ classMapping[item.class] }}
                    </n-tag>
                    <template v-if="item.tags">
                        <n-tag class="mr-5" size="small" v-for="tag in item.tags" :bordered="false" type="info">
                            {{ tag.name }}
                        </n-tag>
                    </template>
                    <n-tag class="mr-5" size="small" v-if="item.isBanner" :bordered="false" type="warning">轮播</n-tag>
                    <n-tag class="mr-5" size="small" v-if="item.isUp" :bordered="false" type="error">置顶</n-tag>
                </div>
                <div class="mt-10 pa-10 abstract">{{ item.abstract == '' ? '本文暂无摘要' : item.abstract }}</div>
                <div class="flex text-small text-gray mt-5">
                    <div class="flex time-item">
                        <div class="mr-10">创建时间</div>
                        <n-time :time="item.creationTime" />
                    </div>
                    <div class="flex time-item justify-center">
                        <div class="mr-10">上次更新</div>
                        <n-time v-if="item.lastUpdateTime" :time="item.lastUpdateTime" />
                        <n-time v-else :time="item.creationTime" />
                    </div>
                    <div class="flex time-item justify-end">
                        <div class="mr-10">发布时间</div>
                        <n-time v-if="item.releaseTime" :time="item.releaseTime" />
                        <div v-else>暂未发布</div>
                    </div>
                </div>
            </div>
        </div>
        <div class="flex justify-center">
            <n-pagination v-model:page="page" :page-count="count" v-model:page-size="number"
                :page-sizes="[10, 20, 40, 80]" show-size-picker />
        </div>
    </div>
</template>

<script>
import { ArrowAutofitUp24Filled, CameraSwitch24Filled, Delete16Regular } from '@vicons/fluent'
import { article } from '../plugins/api'

export default {
    name: "List",
    components: { ArrowAutofitUp24Filled, CameraSwitch24Filled, Delete16Regular },
    data: () => ({
        loading: true,
        keyword: '',
        state: 0,
        number: 20,
        count: 0,
        page: 1,
        list: [],
        classMapping: {},
        tags: [],
    }),
    methods: {
        init() {
            this.initClass();
            this.initTags();
            this.getNumber();
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
        getNumber() {
            this.loading = true;
            article.getNumber(this.keyword, this.state).then(res => {
                if (res.state) {
                    this.page = 1
                    this.count = parseInt((res.data - 1) / this.number) + 1;
                    this.getList()
                } else {
                    window.$message.warning('获取文章数量失败');
                    this.loading = false;
                }
            }).catch(err => {
                this.loading = false;
                window.$message.warning("获取文章数量出错");
            })
        },
        getList() {
            article.getList(this.keyword, this.state, this.page, this.number).then(res => {
                if (res.state) {
                    this.list = res.data
                    this.loading = false;
                } else {
                    window.$message.warning('获取文章列表失败');
                    this.loading = false;
                }
            }).catch(err => {
                this.loading = false;
                window.$message.warning("获取文章列表出错");
            })
        },
        toEdit(id){
            this.$router.push('/edit/'+id);
        }
    },
    mounted() {
        this.init()
    },
};
</script>

<style scoped>
.article-list{
    min-height: calc(100vh - 201px);
}

.article-item {
    cursor: pointer;
}

.article-item:hover {
    background-color: var(--color-dark-3);
}

.title {
    font-size: 16px;
}

.abstract {
    background-color: var(--color-dark-1);
}

.time-item {
    width: 33.33%;
}
</style>
