<template>
    <div class="app-content no-select">
        <div class="flex mb-10">
            <div class="card carousel mr-10">
                <n-carousel direction="vertical" dot-placement="right" show-arrow style="border-radius: 8px;">
                    <div v-if="banner.length == 0" class="card carousel-item welcome flex align-center justify-center">
                        <div>欢迎来访 👏</div>
                    </div>
                    <div v-for="item in banner" class="card carousel-item pa-10" @click="open(item.id)">
                        <div class="banner-title">{{ item.title }}</div>
                        <n-time v-if="item.releaseTime" class="text-gray" :time="item.releaseTime" />
                        <div class="banner-abstract">{{ item.abstract }}</div>
                        <div class="banner-foot flex align-center">
                            <n-tag class="mr-5" size="small" :bordered="false" type="success">
                                {{ classMapping[item.class] }}
                            </n-tag>
                            <template v-if="item.tags">
                                <n-tag class="mr-5" size="small" v-for="tag in item.tags" :bordered="false" type="info">
                                    {{ tag.name }}
                                </n-tag>
                            </template>
                        </div>
                    </div>
                </n-carousel>
            </div>
            <div class="full-width">
                <div class="card readme mb-10 full-width">
                    <div class="flex pa-10 border-bottom">
                        <n-avatar class="mr-10" size="large"
                            src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg" />
                        <div>
                            <div>Skye Zhang</div>
                            <div class="text-gray text-small">@skye-z</div>
                        </div>
                    </div>
                    <div>
                        <n-button class="social-btn" quaternary>
                            <template #icon>
                                <n-icon>
                                    <Github />
                                </n-icon>
                            </template>
                        </n-button>
                        <n-button class="social-btn" quaternary disabled>
                            <template #icon>
                                <n-icon>
                                    <QqOutlined />
                                </n-icon>
                            </template>
                        </n-button>
                        <n-button class="social-btn" quaternary disabled>
                            <template #icon>
                                <n-icon>
                                    <Linkedin />
                                </n-icon>
                            </template>
                        </n-button>
                        <n-button class="social-btn" quaternary disabled>
                            <template #icon>
                                <n-icon>
                                    <Discord />
                                </n-icon>
                            </template>
                        </n-button>
                    </div>
                </div>
                <div class="card tag-cloud pa-10 full-width">标签云</div>
            </div>
        </div>
        <div class="article-box">
            <div v-for="item in tops" class="card article-item mb-10 pa-10" @click="open(item.id)">
                <div class="mb-5 title">{{ item.title }}</div>
                <div class="flex align-center">
                    <n-tag class="mr-5" v-if="item.releaseTime" size="small" :bordered="false">
                        <n-time :time="item.releaseTime" type="relative" />
                    </n-tag>
                    <n-tag class="mr-5" size="small" :bordered="false" type="success">
                        {{ classMapping[item.class] }}
                    </n-tag>
                    <template v-if="item.tags">
                        <n-tag class="mr-5" size="small" v-for="tag in item.tags" :bordered="false" type="info">
                            {{ tag.name }}
                        </n-tag>
                    </template>
                    <n-tag class="mr-5" size="small" :bordered="false" type="error">置顶</n-tag>
                </div>
                <div v-if="item.abstract" class="mt-10 pa-10 abstract">{{ item.abstract }}</div>
            </div>
            <div v-for="item in list" class="card article-item mb-10 pa-10" @click="open(item.id)">
                <div class="mb-5 title">{{ item.title }}</div>
                <div class="flex align-center">
                    <n-tag class="mr-5" v-if="item.releaseTime" size="small" :bordered="false">
                        <n-time :time="item.releaseTime" type="relative" />
                    </n-tag>
                    <n-tag class="mr-5" size="small" :bordered="false" type="success">
                        {{ classMapping[item.class] }}
                    </n-tag>
                    <template v-if="item.tags">
                        <n-tag class="mr-5" size="small" v-for="tag in item.tags" :bordered="false" type="info">
                            {{ tag.name }}
                        </n-tag>
                    </template>
                </div>
                <div v-if="item.abstract" class="mt-10 pa-10 abstract">{{ item.abstract }}</div>
            </div>
            <div class="text-center">
                <n-button class="load-more" strong secondary>加载更多</n-button>
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
        loading: true,
        number: 20,
        page: 1,
        list: [],
        tops: [],
        banner: [],
        classMapping: {},
        tags: [],
    }),
    methods: {
        init() {
            this.initClass();
            this.initTags();
            this.getList(true, false);
            this.getList(false, true);
            this.getList(false, false);
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
        getList(isBanner, isUp) {
            article.getList(isBanner, isUp, 2, this.page, this.number).then(res => {
                if (res.state) {
                    if (isBanner) this.banner = res.data ? res.data : []
                    else if (isUp) this.tops = res.data ? res.data : []
                    else {
                        let list = []
                        for (let i in res.data) {
                            let item = res.data[i]
                            if (item.isBanner || item.isUp) continue
                            list.push(item)
                        }
                        this.list = list
                    }
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
        open(id) {
            this.$router.push('/info/' + id)
        }
    },
    mounted() {
        this.init()
    },
};
</script>

<style scoped>
.welcome {
    font-size: 32px;
    font-weight: bold;
}

.carousel {
    width: 580px;
    height: 300px;
}

.carousel-item {
    width: 580px;
    height: 300px;
    position: relative;
    cursor: pointer;
}

.banner-title {
    font-size: 24px;
    font-weight: bold;
}

.banner-abstract {
    font-size: 18px;
    line-height: 1.5;
    width: calc(100% - 20px);

    display: -webkit-box;
    -webkit-line-clamp: 6;
    -webkit-box-orient: vertical;
    overflow: hidden;
    text-overflow: ellipsis;
}

.banner-foot {
    position: absolute;
    bottom: 10px;
    left: 10px;
}

.readme {
    height: 97px;
}

.tag-cloud {
    height: 192px;
}

.article-list {
    height: 200px;
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

.load-more {
    width: 40%;
}

.readme:deep(.social-btn) {
    width: 25%;
}
</style>
