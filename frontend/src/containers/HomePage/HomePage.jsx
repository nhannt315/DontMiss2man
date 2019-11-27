import React, {useState, useEffect, useRef} from 'react';
import PropTypes from 'prop-types';
import {Row, Col, Select, Pagination} from 'antd';
import scrollToComponent from 'react-scroll-to-component';
import {connect} from 'react-redux';

import './HomePage.scss';
import * as actions from '../../store/actions';
import {SORT_OPTIONS, NUMBER_OF_ITEMS} from '../../constants/common';
import ListHelper from '../../helpers/list_helper';
import i18n from '../../config/i18n';
import BuildingList from '../../components/BuildingList';
import ListPlaceholder from '../../components/ListPlaceholder';


const HomePage = ({list, loading, totalPages, fetchBuildings}) => {
  const firstElement = useRef(null);
  const [page, setPage] = useState(1);
  const [perPage, setPerPage] = useState(NUMBER_OF_ITEMS[0].key);
  useEffect(() => {
    scrollToComponent(firstElement.current);
    fetchBuildings(page, perPage);
  }, [page, perPage]);
  console.log(totalPages, perPage);
  const sortOptionList = ListHelper.generateListFromObject(SORT_OPTIONS);
  return (
    <div className="homepage">
      <div ref={firstElement} />
      <Row>
        <Col span={16}>
          <Row>
            <Col span={8}>
              {i18n.t('homepage.sort')}
              <Select className="sort-filter" defaultValue={SORT_OPTIONS.recommended.key}>
                {sortOptionList.map(item => {
                  return <Select.Option key={item.key} value={item.key}>{item.value}</Select.Option>;
                })}
              </Select>
            </Col>
            <Col span={8}>
              {i18n.t('homepage.number_of_items')}
              <Select value={perPage} className="sort-filter" defaultValue={NUMBER_OF_ITEMS[0].key}
                      onChange={value => setPerPage(value)}>
                {NUMBER_OF_ITEMS.map(item => {
                  return <Select.Option key={item.key} value={item.key}>{item.value}</Select.Option>;
                })}
              </Select>
            </Col>
          </Row>
          <Row />
          <Row>
            <div className="list">
              {loading ? <ListPlaceholder itemCount={perPage} /> : <BuildingList buildingList={list} />}
            </div>
          </Row>
        </Col>
        <Col span={8} />
      </Row>
      <Row>
        <Col className="building-list-pagination" span={16}>
          <Pagination size="small" current={page} pageSize={perPage}
                      onChange={currentPage => setPage(currentPage)}
                      total={totalPages} />
        </Col>
      </Row>
    </div>
  );
};

HomePage.propTypes = {
  list: PropTypes.array,
  loading: PropTypes.bool,
  totalPages: PropTypes.number,
  fetchBuildings: PropTypes.func,
};

HomePage.defaultProps = {
  list: [],
  loading: false,
  totalPages: 1,
  fetchBuildings: () => {
  },
};

const mapStateToProps = state => ({
  list: state.building.list,
  loading: state.building.loading,
  totalPages: state.building.totalPages,
});

const mapDispatchToProps = dispatch => ({
  fetchBuildings: (page, perPage) => dispatch(actions.fetchBuildings(page, perPage)),
});

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
