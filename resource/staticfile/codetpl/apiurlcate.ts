import { defHttp } from '@/utils/http';
enum Api {
    getList = '/modname/filename/get_list',
    getCate = '/modname/filename/get_cate',
    save = '/modname/filename/save',
    upStatus = '/modname/filename/upStatus',
    del = '/modname/filename/del',
}

//列表数据
export function getList(params: object) {
  return defHttp.get({ url: Api.getList, params:params }, { errorMessageMode: 'none' });
}
//列表选项数据
export function getCate(params: object) {
    return defHttp.get({ url: Api.getCate, params:params }, { errorMessageMode: 'none' });
  }
//提交数据
export function save(params: object) {
    return defHttp.post({ url: Api.save, params:params}, { errorMessageMode: 'message' });
}
//更新状态
export function upStatus(params: object) {
    return defHttp.post({ url: Api.upStatus, params:params}, { errorMessageMode: 'message' });
}
//删除数据
export function del(params: object) {
    return defHttp.delete({ url: Api.del, params:params}, { errorMessageMode: 'message' });
}
/**数据类型 */
export interface DataItem {
    id:number,
    name: string;
}