<template>
    <div class="app-content no-select">
        <div class="mb-10 pl-10">
            <div class="flex" v-if="classId">
                <div class="text-gray mr-5">分类: </div>
                <div>{{ classMapping[classId] }}</div>
            </div>
            <div class="flex" v-else-if="keyword">
                <div class="text-gray mr-5">搜索: </div>
                <div>{{ keyword }}</div>
            </div>
        </div>
        <article-list :list="list" />
        <div class="text-center">
            <n-button class="load-more" :disabled="!next" strong secondary @click="getList">加载更多</n-button>
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
        keyword: '',
        number: 0,
        page: 0,
        list: [],
        next: false,
        classMapping: {},
        tags: [],
    }),
    methods: {
        init() {
            this.initData()
            this.loading = true
            let param = this.$route.query
            if (param.class != undefined) {
                this.classId = param.class
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
                this.tops = JSON.parse(cacheTags)
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
            article.getList(true, true, this.classId, 2, this.page, this.number).then(res => {
                if (res.state) {
                    if (res.data == null) res.data = []
                    this.next = res.data.length == this.number
                    if (res.data.length == 0) {
                        window.$message.warning('已经到底啦');
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
</style>
