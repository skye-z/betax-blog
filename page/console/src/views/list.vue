<template>
    <div class="app-content no-select">
        <div class="loading-block" v-if="loading">
            <n-spin />
        </div>
        <n-input-group class="title mb-10">
            <n-input-group-label>关键词</n-input-group-label>
            <n-input v-model:value="keyword" type="text" placeholder="你在找些什么呢?" />
            <n-input-group-label>状态</n-input-group-label>
            <n-select v-model:value="state" label-field="name" value-field="id" style="width: 130px;" :options="states"
                placeholder="筛选状态" />
            <n-button strong secondary type="primary" @click="getNumber">筛选</n-button>
            <n-button strong secondary @click="clean">清除</n-button>
        </n-input-group>
        <div class="card mb-10 article-list">
            <div v-for="(item, index) in list" class="article-item pa-10" @click="toEdit(item.id)"
                :class="{ 'border-top': index != 0 }">
                <n-button-group class="float-right" v-if="item.state != 4">
                    <n-button quaternary circle :type="item.isUp ? 'primary' : 'default'"
                        @click.stop="switchMeta(index, item.isBanner, !item.isUp)">
                        <template #icon>
                            <n-icon>
                                <ArrowAutofitUp24Filled />
                            </n-icon>
                        </template>
                    </n-button>
                    <n-button quaternary circle :type="item.isBanner ? 'primary' : 'default'"
                        @click.stop="switchMeta(index, !item.isBanner, item.isUp)">
                        <template #icon>
                            <n-icon>
                                <CameraSwitch24Filled />
                            </n-icon>
                        </template>
                    </n-button>
                    <n-button v-if="item.state == 2" quaternary circle type="default"
                        @click.stop="switchVisibility(index, 3)">
                        <template #icon>
                            <n-icon>
                                <EyeOff24Filled />
                            </n-icon>
                        </template>
                    </n-button>
                    <n-button v-else-if="item.state == 3" quaternary circle type="default"
                        @click.stop="switchVisibility(index, 2)">
                        <template #icon>
                            <n-icon>
                                <Eye24Filled />
                            </n-icon>
                        </template>
                    </n-button>
                    <n-button quaternary circle @click.stop="remove(index)">
                        <template #icon>
                            <n-icon>
                                <Delete16Regular />
                            </n-icon>
                        </template>
                    </n-button>
                </n-button-group>
                <n-button class="float-right" v-else quaternary circle @click.stop="restore(index)">
                    <template #icon>
                        <n-icon>
                            <ArrowReset24Filled />
                        </n-icon>
                    </template>
                </n-button>
                <div class="mb-5 title">{{ item.title }}</div>
                <div class="flex align-center">
                    <n-tag class="mr-5" size="small" v-if="item.state == 1" :bordered="false">
                        <template #icon>
                            <n-icon>
                                <PresenceAway10Filled />
                            </n-icon>
                        </template>
                        草稿
                    </n-tag>
                    <n-tag class="mr-5" size="small" v-else-if="item.state == 2" type="success" :bordered="false">
                        <template #icon>
                            <n-icon>
                                <PresenceAvailable10Filled />
                            </n-icon>
                        </template>
                        公开
                    </n-tag>
                    <n-tag class="mr-5" size="small" v-else-if="item.state == 3" type="warning" :bordered="false">
                        <template #icon>
                            <n-icon>
                                <PresenceDnd10Filled />
                            </n-icon>
                        </template>
                        私密
                    </n-tag>
                    <n-tag class="mr-5" size="small" v-else-if="item.state == 4" type="error" :bordered="false">
                        <template #icon>
                            <n-icon>
                                <PresenceOffline10Regular />
                            </n-icon>
                        </template>
                        删除
                    </n-tag>
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
import {
    PresenceAway10Filled, PresenceAvailable10Filled, PresenceDnd10Filled,
    PresenceOffline10Regular, ArrowAutofitUp24Filled, CameraSwitch24Filled,
    ArrowReset24Filled, Delete16Regular, Eye24Filled, EyeOff24Filled
} from '@vicons/fluent'
import { article } from '../plugins/api'

export default {
    name: "List",
    components: {
        PresenceAway10Filled, PresenceAvailable10Filled, PresenceDnd10Filled,
        PresenceOffline10Regular, ArrowAutofitUp24Filled, CameraSwitch24Filled,
        ArrowReset24Filled, Delete16Regular, Eye24Filled, EyeOff24Filled
    },
    data: () => ({
        loading: true,
        keyword: '',
        state: 0,
        states: [
            { id: 0, name: '全部' },
            { id: 1, name: '草稿' },
            { id: 2, name: '公开' },
            { id: 3, name: '私密' },
            { id: 4, name: '删除' },
        ],
        number: 20,
        count: 0,
        page: 1,
        list: [],
        stateMapping: {
            1: '草稿',
            2: '公开',
            3: '私密',
            4: '删除',
        },
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
        clean(){
            this.keyword = ''
            this.state = 0
            this.getNumber()
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
        remove(index) {
            let item = this.list[index]
            window.$dialog.warning({
                title: '操作确认',
                content: '文章删除只是将其标记为删除, 你可以通过点击垃圾清理按钮彻底清除, 或者等待每7日自动清理, 确认要删除此文章吗?',
                positiveText: '确认删除',
                negativeText: '取消',
                onPositiveClick: () => {
                    article.remove(item.id).then(res => {
                        if (res.state) {
                            item.state = 4
                            window.$message.success('文章删除成功');
                        }
                        else window.$message.warning('删除文章失败');
                    }).catch(err => {
                        window.$message.warning("发生意料之外的错误");
                    })
                }
            })
        },
        restore(index) {
            let item = this.list[index]
            let state = item.releaseTime ? 3 : 1
            article.switch(item.id, state).then(res => {
                if (res.state) {
                    item.state = state
                    window.$message.success('文章恢复成功');
                }
                else window.$message.warning('文章恢复失败');
            }).catch(err => {
                window.$message.warning("发生意料之外的错误");
            })
        },
        switchVisibility(index, state) {
            let item = this.list[index]
            article.switch(item.id, state).then(res => {
                if (res.state) {
                    item.state = state
                    window.$message.success('操作成功');
                }
                else window.$message.warning('操作失败');
            }).catch(err => {
                window.$message.warning("发生意料之外的错误");
            })
        },
        switchMeta(index, isBanner, isUp) {
            if (isBanner && isUp) {
                window.$message.warning('置顶和轮播只能二选一哦');
                return false
            }
            let item = this.list[index]
            article.meta(item.id, isBanner, isUp).then(res => {
                if (res.state) {
                    item.isBanner = isBanner
                    item.isUp = isUp
                    window.$message.success('操作成功');
                }
                else window.$message.warning('操作失败');
            }).catch(err => {
                window.$message.warning("发生意料之外的错误");
            })
        },
        toEdit(id) {
            this.$router.push('/edit/' + id);
        }
    },
    mounted() {
        this.init()
    },
};
</script>

<style scoped>
.article-list {
    min-height: calc(100vh - 201px);
}

.article-item {
    cursor: pointer;
}

.article-item:hover {
    background-color: var(--color-light-3);
}

.dark .article-item:hover {
    background-color: var(--color-dark-3);
}

.title {
    font-size: 16px;
}

.abstract {
    background-color: var(--color-light-1);
}

.dark .abstract {
    background-color: var(--color-dark-3);
}

.time-item {
    width: 33.33%;
}
</style>
