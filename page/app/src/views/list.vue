<template>
    <div class="app-content no-select">
        <div v-if="loading" class="loading-block"><n-spin /></div>
        <div class="mb-10 pl-10">
            <div class="flex" v-if="classId">
                <div class="text-gray mr-5">分类: </div>
                <div>{{ classMapping[classId] }}</div>
            </div>
            <div class="flex" v-else-if="tagId">
                <div class="text-gray mr-5">标签: </div>
                <div>{{ tagsMapping[tagId] }}</div>
            </div>
            <div class="flex" v-else-if="keyword">
                <div class="text-gray mr-5">搜索: </div>
                <div>{{ keyword }}</div>
            </div>
        </div>
        <article-list :list="list" />
        <div class="text-center" v-if="list.length > 0">
            <n-button class="load-more" :disabled="!next" strong secondary @click="getList">加载更多</n-button>
        </div>
        <div class="none" v-if="!loading && list.length == 0">
            没有相关文章
        </div>
    </div>
</template>

<script>
import ArticleList from '../components/articleList.vue'
import { article } from '../plugins/api'

export default {
    name: "List",
    components: { ArticleList },
    data: () => ({
        loading: true,
        classId: null,
        tagId: null,
        keyword: '',
        number: 0,
        page: 0,
        list: [],
        next: false,
        classMapping: {},
        tagsMapping: [],
    }),
    methods: {
        init() {
            this.initData()
            this.loading = true
            let param = this.$route.query
            if (param.class != undefined) {
                this.classId = param.class
                this.getList()
            } else if (param.tag != undefined) {
                this.tagId = param.tag
                this.getList()
            } else if (param.q != undefined) {
                this.keyword = param.q
                this.searchList()
            }
        },
        initData() {
            this.number = 20
            this.page = 1
            this.next = false
            this.list = []
            let cacheTags = localStorage.getItem('cache:tags');
            if (cacheTags != undefined) {
                cacheTags = JSON.parse(cacheTags)
                for (let i in cacheTags) {
                    this.tagsMapping[cacheTags[i].id] = cacheTags[i].name
                }
            }
            let cacheClass = localStorage.getItem('cache:class');
            if (cacheClass != undefined) {
                cacheClass = JSON.parse(cacheClass)
                for (let i in cacheClass) {
                    this.classMapping[cacheClass[i].id] = cacheClass[i].name
                }
            }
        },
        getList() {
            article.getList(true, true, this.classId, this.tagId, 2, this.page, this.number).then(res => {
                if (res.state) {
                    if (res.data == null) res.data = []
                    this.next = res.data.length == this.number
                    if (res.data.length == 0) {
                        window.$message.warning('已经到底啦');
                        this.loading = false;
                        return false;
                    }
                    for (let i in res.data) {
                        this.list.push(res.data[i])
                    }
                    if (this.next) this.page = this.page + 1;
                } else {
                    window.$message.warning('获取文章列表失败');
                }
                this.loading = false;
            }).catch(err => {
                this.loading = false;
                window.$message.warning("获取文章列表出错");
            })
        },
        searchList() {
            article.search(this.keyword, this.page, this.number).then(res => {
                if (res.state) {
                    this.list = res.data;
                } else {
                    window.$message.warning('获取文章列表失败');
                }
                this.loading = false;
            }).catch(err => {
                this.loading = false;
                window.$message.warning("获取文章列表出错");
            })
        }
    },
    mounted() {
        this.init()
    },
    watch: {
        '$route': function (to, from) {
            if (from.name == to.name) {
                this.init()
            }
        },
    },
};
</script>

<style scoped>
.load-more {
    width: 40%;
}

.none{
    text-align: center;
    padding: calc(50vh - 102px) 0;
    font-size: 32px;
    color: var(--text-color-light-2);
}
</style>
