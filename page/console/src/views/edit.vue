<template>
    <div class="app-content no-select">
        <div class="flex mb-10">
            <n-input-group class="title mr-10">
                <n-input-group-label>标题</n-input-group-label>
                <n-input v-model:value="form.title" :disabled="loading" type="text" placeholder="为文章取个标题吧~" />
            </n-input-group>
            <n-button-group>
                <n-button strong secondary :disabled="loading" @click="save(false)">保存</n-button>
                <n-button strong secondary v-if="form.state == 1" :loading="loading" type="primary"
                    @click="publish">发布文章</n-button>
                <n-button strong secondary v-else-if="form.state == 2" :loading="loading" type="warning"
                    @click="switchState(3)">转为私密</n-button>
                <n-button strong secondary v-else-if="form.state == 3" :loading="loading" type="info"
                    @click="switchState(2)">转为公开</n-button>
            </n-button-group>
        </div>
        <div class="setting flex mb-10">
            <n-input-group class="class mr-10">
                <n-input-group-label>分类</n-input-group-label>
                <n-select v-model:value="form.classId" :disabled="loading" label-field="name" value-field="id"
                    :options="classList" placeholder="请选择分类" />
            </n-input-group>
            <n-input-group class="tags">
                <n-input-group-label>标签</n-input-group-label>
                <n-select class="tags" v-model:value="form.tags" :disabled="loading" label-field="name" value-field="id"
                    multiple :options="tagList" placeholder="请选择标签" />
            </n-input-group>
        </div>
        <n-input-group class="title mb-10">
            <n-input-group-label>摘要</n-input-group-label>
            <n-input v-model:value="form.abstract" :disabled="loading || ai" type="text" placeholder="不如让 AI 生成试试?" />
            <n-button strong secondary :disabled="loading" :loading="ai" @click="aiBuild">
                <template #icon>
                    <n-icon>
                        <BrainCircuit20Filled />
                    </n-icon>
                </template>
                AI摘要
            </n-button>
        </n-input-group>
        <div id="vditor" class="mb-10">
            <div class="tips">
                <n-spin clss="mb-10" :size="78" :stroke-width="24" />
                <div>编辑器初始化中...</div>
            </div>
        </div>
        <div class="text-center text-gray">
            <span>上次保存: </span>
            <n-time v-if="last > 0" :time="last" />
            <span v-else>每10秒自动保存1次哦 OuO</span>
        </div>
    </div>
</template>

<script>
import { BrainCircuit20Filled } from '@vicons/fluent'
import { article } from '../plugins/api'
import Vditor from 'vditor'
import 'vditor/dist/index.css';

export default {
    name: "Info",
    components: { BrainCircuit20Filled },
    data: () => ({
        model: 'add',
        form: {
            id: 0,
            tags: [],
            title: '',
            classId: null,
            isBanner: false,
            isUp: false,
            abstract: '',
            state: 0,
            releaseTime: null
        },
        editor: null,
        cache: "",
        classList: [],
        tagList: [],
        timer: 0,
        isFocused: true,
        last: 0,
        loading: false,
        ai: false
    }),
    methods: {
        init() {
            if (this.$route.params.id) {
                this.form.id = parseInt(this.$route.params.id);
                this.model = 'edit'
            } else {
                this.form.id = 0;
                this.model = 'add'
            }
            this.initClass();
            this.initTags();
            this.initData();
        },
        initClass() {
            if (this.classList.length > 0) return false;
            article.getClass().then(res => {
                if (res.state) this.classList = res.data
                else {
                    window.$message.warning('获取分类列表失败');
                }
            }).catch(err => {
                window.$message.warning("发生意料之外的错误");
            })
        },
        initTags() {
            if (this.tagList.length > 0) return false;
            article.getTags().then(res => {
                if (res.state) this.tagList = res.data
                else {
                    window.$message.warning('获取标签列表失败');
                }
            }).catch(err => {
                window.$message.warning("发生意料之外的错误");
            })
        },
        initEditor(content) {
            if (this.editor != null) {
                this.editor.setValue(content)
                this.controlTimer(true)
                return false;
            }
            let token = localStorage.getItem("access:token");
            this.editor = new Vditor('vditor', {
                theme: 'dark',
                typewriterMode: true,
                toolbarConfig: {
                    pin: true,
                },
                toolbar: ['undo', 'redo', '|', 'headings', 'bold', 'italic', 'strike', '|', 'line', 'quote', 'list', 'ordered-list', 'check', 'outdent', 'indent', 'code', 'link', 'upload', '|', 'both', 'preview', 'outline'],
                cache: {
                    enable: false,
                },
                preview: {
                    markdown: {
                        codeBlockPreview: false,
                        mathBlockPreview: false
                    }
                },
                upload: {
                    url: '/api/upload',
                    headers: {
                        Authorization: 'Bearer ' + token
                    },
                    fieldName: 'files',
                    accept: 'image/*'
                },
                after: () => {
                    this.editor.setValue(content)
                }
            })
        },
        initData() {
            if (this.form.id == 0) {
                this.initEditor('')
                return false;
            }
            article.getInfo(this.form.id).then(res => {
                if (res.state) {
                    let info = res.data;
                    this.form.title = info.title;
                    this.form.classId = info.class;
                    this.form.isBanner = info.isBanner;
                    this.form.isUp = info.isUp;
                    this.form.abstract = info.abstract;
                    this.form.state = info.state;
                    this.form.abstract = info.abstract;
                    this.form.releaseTime = info.releaseTime;
                    let tags = [];
                    if (info.tags) {
                        for (let i in info.tags) {
                            tags.push(info.tags[i].id)
                        }
                    }
                    this.form.tags = tags;

                    this.cache = JSON.stringify(this.form) + "|" + info.content;
                    this.initEditor(info.content);
                } else this.initEditor('');
            }).catch(err => {
                console.log(err)
                this.initEditor('')
            })
        },
        aiBuild() {
            this.save(true)
            this.ai = true
            article.aiBuild(this.form.id).then(res => {
                if (res.state) this.form.abstract = res.data
                else window.$message.warning(res.message ? res.message : 'AI摘要失败')
                this.ai = false
            }).catch(err => {
                console.log(err)
                this.ai = false
                window.$message.warning('AI摘要出错')
            })
        },
        save(auto) {
            if (this.editor == null) return false;
            let content = this.editor.getValue();
            if (auto) {
                // 检查是否有更新
                let check = JSON.stringify(this.form);
                check += "|" + content
                if (this.cache == check) return false;
                else this.cache = check
            }
            // 检查内容
            if (content.trim().length == 0) return false;
            // 检查标题
            let title = this.form.title;
            if (title.trim().length == 0) title = '草稿 #' + new Date().getTime()
            // 检查分类
            let classId = this.form.classId;
            if (classId == null) classId = 0;
            if (this.model == 'add') this.addArticle({
                title, content, classId,
                isBanner: this.form.isBanner,
                isUp: this.form.isUp,
                abstract: this.form.abstract
            }, auto);
            else this.updateArticle({
                title, content, classId,
                id: this.form.id,
                isBanner: this.form.isBanner,
                isUp: this.form.isUp,
                abstract: this.form.abstract,
                tags: this.form.tags,
                state: this.form.state,
                releaseTime: this.form.releaseTime
            }, auto);
            this.last = new Date().getTime();
        },
        addArticle(form, auto) {
            article.add(form.isBanner, form.isUp, 1, form.title, form.abstract, form.classId, form.tags, form.content, 1, null).then(res => {
                if (res.state) {
                    if (!auto) window.$message.success('保存成功')
                    this.$router.replace('/edit/' + res.data)
                    this.form.id = res.data
                    this.model = 'edit'
                } else {
                    if (!auto) window.$message.warning(res.message ? res.message : '保存失败')
                }
            }).catch(err => {
                console.log(err)
                window.$message.error('保存出错')
            })
        },
        updateArticle(form, auto) {
            article.edit(form.id, form.isBanner, form.isUp, 1, form.title, form.abstract, form.classId, form.tags, form.content, form.state, form.releaseTime).then(res => {
                if (!auto) {
                    if (res.state) window.$message.success('保存成功')
                    else window.$message.warning(res.message ? res.message : '保存失败')
                }
            }).catch(err => {
                console.log(err)
                window.$message.error('保存出错')
            })
        },
        publish() {
            this.loading = true
            let time = this.form.id == 0 ? 1000 : 100
            if (this.form.id == 0) this.save(false)
            setTimeout(() => {
                article.publish(this.form.id).then(res => {
                    if (res.state) {
                        window.$message.success('发布成功')
                        this.form.state = 2
                    } else window.$message.warning(res.message ? res.message : '发布失败')
                    this.loading = false
                }).catch(err => {
                    console.log(err)
                    window.$message.error('发布出错')
                    this.loading = false
                })
            }, time)
        },
        switchState(state) {
            this.loading = true
            article.switch(this.form.id, state).then(res => {
                if (res.state) {
                    window.$message.success('操作成功')
                    this.form.state = state
                } else window.$message.warning(res.message ? res.message : '操作失败')
                this.loading = false
            }).catch(err => {
                console.log(err)
                window.$message.error('操作出错')
                this.loading = false
            })
        },
        controlTimer(state) {
            if (state) {
                if (this.isFocused) {
                    if (this.timer != 0) clearInterval(this.timer);
                    this.timer = setInterval(() => {
                        this.save(true)
                    }, 10000);
                }
            } else clearInterval(this.timer);
        },
        handleFocus() {
            this.isFocused = true;
            this.controlTimer(true);
        },
        handleBlur() {
            this.isFocused = false;
            this.controlTimer(false);
        }
    },
    mounted() {
        this.init();
        this.controlTimer(true);
        window.addEventListener('focus', this.handleFocus);
        window.addEventListener('blur', this.handleBlur);
    },
    unmounted() {
        window.removeEventListener('focus', this.handleFocus);
        window.removeEventListener('blur', this.handleBlur);
        this.controlTimer(false);
    },
    beforeRouteLeave(_to, _from, next) {
        console.log('beforeRouteLeave')
        this.controlTimer(false);
        window.removeEventListener('focus', this.handleFocus);
        window.removeEventListener('blur', this.handleBlur);
        next();
    },
    watch: {
        '$route': function (to, from) {
            if (from.name == to.name) {
                this.controlTimer(false);
                this.init()
            }
        },
    },
};
</script>

<style scoped>
.title {
    width: 100%;
}

.class {
    width: 300px;
}

#vditor {
    height: calc(100vh - 275px) !important;
}

#vditor .tips {
    text-align: center;
    font-size: 24px;
    padding: 100px;
}

.tags:deep(.n-tag) {
    background-color: var(--color-dark-3);
}

@media only screen and (max-width: 767px) {
    .setting {
        display: block !important;
    }

    .class {
        width: 100%;
        margin-bottom: 10px;
    }

    #vditor {
        height: calc(100vh - 319px) !important;
    }
}
</style>
