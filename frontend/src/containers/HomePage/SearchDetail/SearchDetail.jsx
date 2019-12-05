import React, {useState} from 'react';
import PropTypes from 'prop-types';
import {Select, Checkbox} from 'antd';
import i18n from '../../../config/i18n';
import './SearchDetail.scss';
import {
  LOWER_RENT_FEE_OPTIONS,
  UPPER_RENT_FEE_OPTIONS,
  LAYOUT_TYPE_OPTIONS,
  UPPER_SIZE_OPTIONS,
  LOWER_SIZE_OPTIONS,
  YEAR_OPTIONS,
} from '../../../constants/searchFilter';

const {Option} = Select;

const SearchDetail = ({searchWithCondition, initialCondition}) => {
  const initialState = {
    rentFee: {
      upper: null,
      lower: null,
      noManagementFee: false,
      noReikin: false,
      noShikikin: false,
      options: [],
    },
    layoutType: [],
    buildingType: [],
    size: {
      upper: null,
      lower: null,
    },
    years_built: null,
  };
  const [searchCondition, setCondition] = useState(initialCondition || initialState);

  return (
    <div className="condition-box">
      <div className="condition-box-title">
        {i18n.t('searchFilter.filterSearch')}
      </div>
      <div className="condition-box-content-wrapper">
        <div className="condition-box-content">
          <dl>
            <dt>{i18n.t('searchFilter.rent_fee')}</dt>
            <dd className="select-area-wrapper">
              <div className="select-area">
                <Select defaultValue={LOWER_RENT_FEE_OPTIONS[0].key} className="select-items"
                        value={searchCondition.rentFee.lower}
                        onChange={value => setCondition({
                          ...searchCondition,
                          rentFee: {...searchCondition.rentFee, lower: value},
                        })}
                >
                  {LOWER_RENT_FEE_OPTIONS.map(ele => (
                    <Option value={ele.key} key={ele.key}>{ele.value}</Option>
                  ))}
                </Select>
                <span className="select-label">〜</span>
                <Select defaultValue={UPPER_RENT_FEE_OPTIONS[0].key} className="select-items"
                        value={searchCondition.rentFee.upper}
                        onChange={value => setCondition({
                          ...searchCondition,
                          rentFee: {...searchCondition.rentFee, upper: value},
                        })}
                >
                  {UPPER_RENT_FEE_OPTIONS.map(ele => (
                    <Option value={ele.key} key={ele.key}>{ele.value}</Option>
                  ))}
                </Select>
              </div>
              <div className="checkboxs">
                <Checkbox.Group
                  value={searchCondition.rentFee.options}
                  onChange={values => {
                    const condition = {};
                    condition.noManagementFee = values.includes('fee');
                    condition.noReikin = values.includes('reikin');
                    condition.noShikikin = values.includes('shikikin');
                    condition.options = values;
                    setCondition({...searchCondition, rentFee: {...searchCondition.rentFee, ...condition}});
                  }}>
                  <div><Checkbox value="fee">{i18n.t('searchFilter.include_management_fee')}</Checkbox></div>
                  <div><Checkbox value="reikin">{i18n.t('searchFilter.no_reikin')}</Checkbox></div>
                  <div><Checkbox value="shikikin">{i18n.t('searchFilter.no_shikikin')}</Checkbox></div>
                </Checkbox.Group>
              </div>
            </dd>
          </dl>
          <dl>
            <dt>{i18n.t('searchFilter.layout_type')}</dt>
            <dd className="layout-detail-list">
              <Checkbox.Group
                value={searchCondition.layoutType}
                onChange={values => setCondition({...searchCondition, layoutType: values})}
              >
                {LAYOUT_TYPE_OPTIONS.map(ele => (
                  <div key={ele} className="checkbox-wrapper"><Checkbox value={ele}>{ele}</Checkbox></div>
                ))}
              </Checkbox.Group>
            </dd>
          </dl>
          <dl>
            <dt>{i18n.t('searchFilter.building_type')}</dt>
            <dd className="building-type-list">
              <Checkbox.Group
                value={searchCondition.buildingType}
                onChange={values => setCondition({...searchCondition, buildingType: values})}>
                <div><Checkbox value="マンション">{i18n.t('searchFilter.mansion')}</Checkbox></div>
                <div><Checkbox value="アパート">{i18n.t('searchFilter.apartment')}</Checkbox></div>
              </Checkbox.Group>
            </dd>
          </dl>
          <dl>
            <dt>{i18n.t('searchFilter.size')}</dt>
            <dd className="select-area-wrapper">
              <div className="select-area">
                <Select defaultValue={LOWER_SIZE_OPTIONS[0].key} className="select-items"
                        value={searchCondition.size.lower}
                        onChange={value => {
                          setCondition({...searchCondition, size: {...searchCondition.size, lower: value}});
                        }}
                >
                  {LOWER_SIZE_OPTIONS.map(ele => (
                    <Option value={ele.key} key={ele.key}>{ele.value}</Option>
                  ))}
                </Select>
                <span className="select-label">〜</span>
                <Select defaultValue={UPPER_SIZE_OPTIONS[0].key} className="select-items"
                        value={searchCondition.size.upper}
                        onChange={value => {
                          setCondition({...searchCondition, size: {...searchCondition.size, upper: value}});
                        }}
                >
                  {UPPER_SIZE_OPTIONS.map(ele => (
                    <Option value={ele.key} key={ele.key}>{ele.value}</Option>
                  ))}
                </Select>
              </div>
            </dd>
          </dl>
          <dl>
            <dt>{i18n.t('searchFilter.years_built')}</dt>
            <dd className="year-select">
              <Select defaultValue={YEAR_OPTIONS[YEAR_OPTIONS.length - 1].key} value={searchCondition.years_built}
                      onChange={value => setCondition({...searchCondition, years_built: value})}
              >
                {YEAR_OPTIONS.map(ele => (
                  <Option key={ele.key} value={ele.key}>{ele.value}</Option>
                ))}
              </Select>
            </dd>
          </dl>
        </div>
        <div className="submit-area" onClick={() => searchWithCondition(searchCondition)}>
          {i18n.t('searchFilter.search_with_condition')}
        </div>
      </div>
    </div>
  );
};

SearchDetail.propTypes = {
  searchWithCondition: PropTypes.func,
  initialCondition: PropTypes.object,
};

SearchDetail.defaultProps = {
  searchWithCondition: () => {
  },
  initialCondition: null,
};

export default SearchDetail;
