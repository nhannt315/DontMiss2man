import React, { useMemo } from 'react';
import useTranslation from 'next-translate/useTranslation';
import { Constants } from 'src/constants/constants';
import Select from 'src/components/Select';
import { Checkbox } from 'src/components/Checkbox';

interface IProps {
  className?: string;
}

const SearchDetail: React.FC<IProps> = ({ className }) => {
  const { t, lang } = useTranslation('searchFilter');
  const constantCreator = useMemo(() => {
    return new Constants(t, lang);
  }, [t, lang]);

  const lowFeeOptions = useMemo(() => constantCreator.getLowFeeOptions(), [
    constantCreator,
  ]);
  const upperFeeOptions = useMemo(() => constantCreator.getUpperFeeOptions(), [
    constantCreator,
  ]);

  const layoutOptions = useMemo(() => constantCreator.getLayoutOptions(), [
    constantCreator,
  ]);

  const upperSizeOptions = useMemo(() => constantCreator.getUpperFeeOptions(), [
    constantCreator,
  ]);
  const lowerSizeOptions = useMemo(
    () => constantCreator.getLowerSizeOptions(),
    [constantCreator]
  );

  const yearOptions = useMemo(() => constantCreator.getYearOptions(), [
    constantCreator,
  ]);

  return (
    <div className={`${className} flex flex-col mx-8`}>
      <div className="bg-gray-100">
        <div className="text-white bg-gray-600 pl-2 py-2 text-sm font-bold">
          {t('filterSearch')}
        </div>
        <div className="pt-2 mx-2 px-2">
          <div className="font-bold text-gray-500 text-sm">{t('rent_fee')}</div>
          <div className="flex flex-row items-center justify-center pt-1">
            <Select className="flex-1" options={lowFeeOptions} />
            <span className="mx-4">〜</span>
            <Select className="flex-1" options={upperFeeOptions} />
          </div>
          <div className="pt-2">
            <div>
              <Checkbox label={t('include_management_fee')} />
            </div>
            <div>
              <Checkbox label={t('no_reikin')} />
            </div>
            <div>
              <Checkbox label={t('no_shikikin')} />
            </div>
          </div>

          <div className="font-bold text-gray-500 text-sm pt-2">
            {t('layout_type')}
          </div>
          <div>
            {layoutOptions.map((option) => (
              <Checkbox className="w-5/12" key={option} label={option} />
            ))}
          </div>

          <div className="font-bold text-gray-500 text-sm pt-2">
            {t('room_facilities')}
          </div>
          <div>
            <Checkbox label={t('with_furniture')} />
          </div>

          <div className="font-bold text-gray-500 text-sm pt-2">
            {t('size')}
          </div>
          <div className="flex flex-row items-center justify-center pt-1">
            <Select className="flex-1" options={lowerSizeOptions} />
            <span className="mx-4">〜</span>
            <Select className="flex-1" options={upperSizeOptions} />
          </div>

          <div className="font-bold text-gray-500 text-sm pt-4">
            {t('years_built')}
          </div>
          <div className="pt-1">
            <div className="w-5/12">
              <Select className="flex-1 w-full" options={yearOptions} />
            </div>
          </div>
        </div>
        <div className="mt-8 text-center bg-blue-300 hover:bg-blue-500 hover:underline text-white py-2 text-base cursor-pointer">
          {t('search_with_condition')}
        </div>
      </div>
    </div>
  );
};

export default SearchDetail;
