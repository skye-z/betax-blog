import request from './request'

function postJSON(url, data) {
    return request({
        url: url,
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: JSON.stringify(data)
    })
}

function post(url, data) {
    let cache = new URLSearchParams();
    for (let key in data) {
        cache.append(key, data[key]);
    }

    return request({
        url: url,
        method: 'POST',
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        },
        data: cache
    })
}

function get(url) {
    return request({
        url: url,
        method: 'GET'
    })
}

export const setting = {
    getPing: () => get('/ping')
}

export const article = {
    add: (isBanner, isUp, type, title, abstract, classId, tags, content, state, releaseTime) => postJSON('/article/add', {
        isBanner, isUp, type, title, abstract, 'class': classId, tagIds: tags, content, state, releaseTime
    }),
    edit: (id, isBanner, isUp, type, title, abstract, classId, tags, content, state, releaseTime) => postJSON('/article/edit', {
        id, isBanner, isUp, type, title, abstract, 'class': classId, tagIds: tags, content, state, releaseTime
    }),
    getInfo: id => get('/article/details/' + id),
    getNumber: (keyword, state) => post('/article/number', { keyword, state }),
    getList: (keyword, state, page, number) => post('/article/list', { isBanner: 1, isUp: 1,keyword, state, page, number }),
    publish: id => post('/article/publish', { id }),
    switch: (id,state) => post('/article/switch', { id,state }),
    remove: id => post('/article/remove', { id }),
    aiBuild: id => get('/article/abstract?id=' + id),
    getClass: () => get('/class'),
    getTags: () => get('/tags')
}