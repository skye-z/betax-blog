import request from './request'

function post(url, data) {
    return request({
        url: url,
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: JSON.stringify(data)
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
    add: (isBanner, isUp, type, title, abstract, classId, content, state, releaseTime) => post('/article/add', {
        isBanner, isUp, type, title, abstract, 'class': classId, content, state, releaseTime
    }),
    edit: (id, isBanner, isUp, type, title, abstract, classId, content, state, releaseTime) => post('/article/edit', {
        id, isBanner, isUp, type, title, abstract, 'class': classId, content, state, releaseTime
    }),
    getInfo: id => get('/article/details/' + id),
    getClass: () => get('/class'),
    getTags: () => get('/tags')
}