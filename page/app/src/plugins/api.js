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

export const common = {
    init: () => get('/init'),
    install: form => postJSON('/install', form),
    getUser: () => get('/user'),
}

export const article = {
    getInfo: id => get('/article/details/' + id),
    search: (keyword, page, number) => post('/article/search', { keyword, page, number }),
    getNumber: (keyword, state) => post('/article/number', { keyword, state }),
    getList: (isBanner, isUp, classId, state, page, number) => post('/article/list', { isBanner: isBanner ? 1:0, isUp: isUp ? 1:0, 'class':classId, state, page, number }),
    getAnyList: () => get('/article/any'),
    getClass: () => get('/class'),
    getTags: () => get('/tags')
}