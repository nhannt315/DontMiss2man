import i18n from '../config/i18n';

export const createSortOptions = () => {
  return {
    recommended: {
      key: 'recommended',
      value: i18n.t('homepage.sort_options.recommended'),
    },
    cheapest: {
      key: 'cheapest',
      value: i18n.t('homepage.sort_options.cheapest_fee'),
    },
    most_expensive: {
      key: 'most_expensive',
      value: i18n.t('homepage.sort_options.most_expensive_fee'),
    },
    newest_building: {
      key: 'newly_built',
      value: i18n.t('homepage.sort_options.newest_building'),
    },
    most_largest: {
      key: 'largest',
      value: i18n.t('homepage.sort_options.most_largest'),
    },
    nearest: {
      key: 'nearest',
      value: i18n.t('homepage.sort_options.nearest'),
    },
  };
};

export const SORT_OPTIONS = createSortOptions();

export const createNumberOfItemOptions = () => {
  return [
    {
      key: 10,
      value: i18n.t('homepage.number_items', {count: 10}),
    },
    {
      key: 15,
      value: i18n.t('homepage.number_items', {count: 15}),
    },
    {
      key: 20,
      value: i18n.t('homepage.number_items', {count: 20}),
    },
    {
      key: 30,
      value: i18n.t('homepage.number_items', {count: 30}),
    },
  ];
};

export const NUMBER_OF_ITEMS = createNumberOfItemOptions();
