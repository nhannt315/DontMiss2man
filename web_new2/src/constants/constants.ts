import { Translate } from 'next-translate';
import { convertYen } from 'src/utils/currency';
import dayjs from 'dayjs';

interface SelectOption {
  label: string | number;
  value: string | number;
}

export class Constants {
  private mLang: string;
  private mTranslate: Translate;
  constructor(t: Translate, currentLang: string) {
    this.mLang = currentLang;
    this.mTranslate = t;
  }

  public getLowFeeOptions(): SelectOption[] {
    const result: SelectOption[] = RENT_FEE.map((ele) => ({
      value: ele * 10000,
      label: convertYen(ele * 10000, this.mLang, this.mTranslate),
    }));
    result.unshift({
      value: 0,
      label: this.mTranslate('searchFilter:no_lower_limit'),
    });
    return result;
  }

  public getUpperFeeOptions(): SelectOption[] {
    const result = RENT_FEE.map((ele) => ({
      value: ele * 10000,
      label: convertYen(ele * 10000, this.mLang, this.mTranslate),
    }));
    result.unshift({
      value: 0,
      label: this.mTranslate('searchFilter:without_limit'),
    });
    return result;
  }

  public getLayoutOptions(): string[] {
    return [
      this.mTranslate('searchFilter:one_room'),
      '1K',
      '1DK',
      '1LDK',
      '2K',
      '2DK',
      '2LDK',
      '3K',
      '3DK',
      '3LDK',
      '4K',
      '4DK',
      '4LDK',
    ];
  }

  public getUpperSizeOptions(): SelectOption[] {
    const result = SIZE_OPTIONS.map((ele) => ({
      value: ele,
      label: `${ele}m2`,
    }));
    result.unshift({
      value: 0,
      label: this.mTranslate('searchFilter:without_limit'),
    });
    return result;
  }

  public getLowerSizeOptions(): SelectOption[] {
    const result = SIZE_OPTIONS.map((ele) => ({
      value: ele,
      label: `${ele}m2`,
    }));
    result.unshift({
      value: 0,
      label: this.mTranslate('searchFilter:no_lower_limit'),
    });
    return result;
  }

  public getYearOptions(): SelectOption[] {
    const result = YEAR_OPTION_LIST.map((ele) => ({
      value: dayjs().year() - ele,
      label: this.mTranslate('searchFilter:within_years', {
        postProcess: 'interval',
        years: ele,
      }),
    }));
    result.push({
      value: 0,
      label: this.mTranslate('searchFilter:not_specified'),
    });
    return result;
  }
}

const YEAR_OPTION_LIST = [1, 3, 5, 7, 10, 15, 20, 25, 30];

const SIZE_OPTIONS = [20, 25, 30, 35, 40, 45, 50, 55, 65, 70, 80, 90, 100];

const RENT_FEE = [
  3,
  3.5,
  4,
  4.5,
  5,
  5.5,
  6,
  6.5,
  7,
  7.5,
  8,
  8.5,
  9,
  9.5,
  10,
  10.5,
  11,
  11.5,
  12,
  12.5,
  13,
  13.5,
  14,
  14.5,
  15,
  15.5,
  16,
  16.5,
  17,
  17.5,
  18,
  18.5,
  19,
  19.5,
  20,
  21,
  22,
  23,
  24,
  25,
  26,
  27,
  28,
  29,
  30,
  35,
  40,
  50,
  100,
];
