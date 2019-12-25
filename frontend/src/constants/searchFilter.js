import moment from 'moment';
import CommonHelper from '../helpers/common';
import i18n from '../config/i18n';

const RENT_FEE = [3, 3.5, 4, 4.5, 5, 5.5, 6, 6.5, 7, 7.5, 8, 8.5,
  9, 9.5, 10, 10.5, 11, 11.5, 12, 12.5, 13, 13.5, 14, 14.5, 15, 15.5, 16, 16.5, 17,
  17.5, 18, 18.5, 19, 19.5, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 35, 40, 50, 100];

export const createUpperFeeOptions = () => {
  const result = RENT_FEE.map(ele => ({key: ele * 10000, value: CommonHelper.convertYen(ele * 10000)}));
  result.unshift({key: null, value: i18n.t('searchFilter.without_limit')});
  return result;
};

export const createLowerFeeOptions = () => {
  const result = RENT_FEE.map(ele => ({key: ele * 10000, value: CommonHelper.convertYen(ele * 10000)}));
  result.unshift({key: null, value: i18n.t('searchFilter.no_lower_limit')});
  return result;
};

export const UPPER_RENT_FEE_OPTIONS = createUpperFeeOptions();

export const LOWER_RENT_FEE_OPTIONS = createLowerFeeOptions();

export const LAYOUT_TYPE_OPTIONS = ['ワンルーム', '1K', '1DK', '1LDK', '2K', '2DK', '2LDK', '3K', '3DK', '3LDK', '4K', '4DK', '4LDK'];

const SIZE_OPTIONS = [20, 25, 30, 35, 40, 45, 50, 55, 65, 70, 80, 90, 100];

export const createUpperSizeOptions = () => {
  const result = SIZE_OPTIONS.map(ele => ({key: ele, value: `${ele}m2`}));
  result.unshift({key: null, value: i18n.t('searchFilter.without_limit')});
  return result;
};

export const createLowerSizeOptions = () => {
  const result = SIZE_OPTIONS.map(ele => ({key: ele, value: `${ele}m2`}));
  result.unshift({key: null, value: i18n.t('searchFilter.no_lower_limit')});
  return result;
};

export const UPPER_SIZE_OPTIONS = createUpperSizeOptions();
export const LOWER_SIZE_OPTIONS = createLowerSizeOptions();

const YEAR_OPTION_LIST = [1, 3, 5, 7, 10, 15, 20, 25, 30];
export const createYearOptions = () => {
  const result = YEAR_OPTION_LIST.map(ele => ({
    key: moment().year() - ele,
    value: i18n.t('searchFilter.within_years', {postProcess: 'interval', years: ele}),
  }));
  result.push({key: null, value: i18n.t('searchFilter.not_specified')});
  return result;
};

export const YEAR_OPTIONS = createYearOptions();

