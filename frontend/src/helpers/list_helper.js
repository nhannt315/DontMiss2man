import {find} from 'lodash';

export default class ListHelper {
  static findObjectValue(list, key){
    const result = find(list, {key});
    if (result)
      return result.value;
    return null;
  }

  static findObject(list, key) {
    return find(list, {key});
  }

  static generateListFromObject(src) {
    return Object.keys(src).map(key => src[key]);
  }
}
