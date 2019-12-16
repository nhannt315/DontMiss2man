import moment from 'moment';
import qs from 'query-string';
import i18n from '../config/i18n';

export default class CommonHelper {
  static convertYen(src) {
    if (!src)
      return '-';
    let result = '';
    if (src > 10000) {
      result = `${this.round(src / 10000, 1)}万円`;
    } else {
      result = `${src}円`;
    }
    return result;
  }


  static round(value, precision) {
    // eslint-disable-next-line no-restricted-properties
    const multiplier = Math.pow(10, precision || 0);
    return Math.round(value * multiplier) / multiplier;
  }

  static getYearBuiltInJap(src) {
    const years = moment().year() - moment(src).year();
    if (years > 2) {
      return i18n.t('common.number_of_year_built', {years});
    } 
      return i18n.t('common.newly_built');
    
  }

  static checkEmptyObject(obj) {
    if (!obj)
      return true;
    return Object.keys(obj).length === 0 && obj.constructor === Object
  }

  static checkLocalstorageStr(src) {
    return ['undefined', 'null'].includes(src);
  }

  static getValueFromQuery(location, key) {
    return qs.parse(location.search, { ignoreQueryPrefix: true })[key];
  }
}
