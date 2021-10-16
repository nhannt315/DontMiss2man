import { Translate } from 'next-translate';

export const convertYen = (
  src: number,
  currentLng: string,
  t: Translate
): string => {
  if (!src) return '-';
  let result = '';
  if (currentLng !== 'ja') {
    return t('common:money_yen', { value: src });
  }
  if (src > 10000) {
    const value = round(src / 10000, 1);
    result = t('common:money_ten_thousand_yen', { value });
  } else {
    result = t('common:money_yen', { value: src });
  }
  return result;
};

export const round = (value: number, precision: number): number => {
  const multiplier = Math.pow(10, precision || 0);
  return Math.round(value * multiplier) / multiplier;
};
