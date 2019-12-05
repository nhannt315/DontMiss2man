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
import SearchDetail from './SearchDetail';

const HomePage = ({list, loading, totalCount, fetchBuildings, history, conditionRedux, sortRedux, currentPage, perPageRedux}) => {
  const firstElement = useRef(null);
  const [isInitialized, setInitialize] = useState(false);
  const [searchCondition, setCondition] = useState(conditionRedux);
  const [page, setPage] = useState(currentPage);
  const [perPage, setPerPage] = useState(perPageRedux || NUMBER_OF_ITEMS[0].key);
  const [sortOption, setSortOption] = useState(sortRedux || SORT_OPTIONS.recommended.key);
  useEffect(() => {
    scrollToComponent(firstElement.current);
    if (isInitialized) {
      fetchBuildings(page, perPage, sortOption, searchCondition);
    }
    if (!isInitialized && list.length === 0) {
      fetchBuildings(page, perPage, sortOption, searchCondition);
      setInitialize(true);
    }
    if (!isInitialized && list.length > 0) {
      setInitialize(true);
    }
  }, [page, fetchBuildings, sortOption, searchCondition, perPage]);


  const searchWithCondition = condition => setCondition(condition);
  const sortOptionList = ListHelper.generateListFromObject(SORT_OPTIONS);
  return (
    <div className="homepage">
      <div ref={firstElement} />
      <Row>
        <Col span={16}>
          <Row>
            <Col span={8}>
              {i18n.t('homepage.sort')}
              <Select className="sort-filter" defaultValue={SORT_OPTIONS.recommended.key}
                      onChange={value => setSortOption(value)}
              >
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
              {loading ? <ListPlaceholder itemCount={perPage} /> :
                <BuildingList history={history} buildingList={list} />}
            </div>
          </Row>
        </Col>
        <Col span={7} offset={1}>
          <SearchDetail searchWithCondition={searchWithCondition} initialCondition={conditionRedux} />
        </Col>
      </Row>
      <Row>
        <Col className="building-list-pagination" span={16}>
          <Pagination size="small" current={page} pageSize={perPage}
                      onChange={cuPage => setPage(cuPage)}
                      total={totalCount} />
        </Col>
      </Row>
    </div>
  );
};

HomePage.propTypes = {
  list: PropTypes.array,
  loading: PropTypes.bool,
  totalCount: PropTypes.number,
  fetchBuildings: PropTypes.func,
  history: PropTypes.object,
  conditionRedux: PropTypes.object,
  sortRedux: PropTypes.string,
  currentPage: PropTypes.number,
  perPageRedux: PropTypes.number,
};

HomePage.defaultProps = {
  list: [],
  loading: false,
  totalCount: 1,
  history: {},
  conditionRedux: null,
  sortRedux: null,
  fetchBuildings: () => {
  },
  currentPage: 1,
  perPageRedux: null,
};

const mapStateToProps = state => ({
  list: state.building.list,
  loading: state.building.loading,
  totalCount: state.building.totalCount,
  conditionRedux: state.building.condition,
  sortRedux: state.building.sort,
  currentPage: state.building.currentPage,
  perPageRedux: state.building.perPage,
});

const mapDispatchToProps = dispatch => ({
  fetchBuildings: (page, perPage, sortOption, condition = null) => dispatch(actions.fetchBuildings(page, perPage, sortOption, condition)),
});

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
