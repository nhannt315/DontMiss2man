import moment from 'moment';
import qs from 'query-string';
import i18n from '../config/i18n';
import {LOCALE_KEY} from '../constants/key';

export default class CommonHelper {
  static convertYen(src) {
    if (!src)
      return '-';
    let result = '';
    const currentLng = localStorage.getItem(LOCALE_KEY) || 'ja';
    if (currentLng !== 'ja') {
      return i18n.t('common.money_yen', {value: src});
    }
    if (src > 10000) {
      const value = this.round(src / 10000, 1);
      result = i18n.t('common.money_ten_thousand_yen', {value});
    } else {
      result = i18n.t('common.money_yen', {value: src});
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
      return i18n.t('common.number_of_year_built', {postProcess: 'interval', years});
    }
    return i18n.t('common.newly_built');

  }

  static checkEmptyObject(obj) {
    if (!obj)
      return true;
    return Object.keys(obj).length === 0 && obj.constructor === Object;
  }

  static checkLocalstorageStr(src) {
    return ['undefined', 'null'].includes(src);
  }

  static getValueFromQuery(location, key) {
    const resultSrc = qs.parse(location.search, {ignoreQueryPrefix: true})[key];
    return parseInt(resultSrc, 10);
  }
}
