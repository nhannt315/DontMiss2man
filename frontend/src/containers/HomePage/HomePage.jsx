import React from 'react';
import PropTypes from 'prop-types';
import {Row, Col, Select, Pagination} from 'antd';
import scrollToComponent from 'react-scroll-to-component';
import {connect} from 'react-redux';

import './HomePage.scss';
import * as actions from '../../store/actions';
import {SORT_OPTIONS, NUMBER_OF_ITEMS} from '../../constants/common';
import ListHelper from '../../helpers/list_helper';
import i18n from '../../config/i18n';
import Layout from '../../components/Layout';
import BuildingList from '../../components/BuildingList';
import ListPlaceholder from '../../components/ListPlaceholder';
import SearchDetail from './SearchDetail';
import CommonHelper from '../../helpers/common';

class HomePage extends React.PureComponent {
  constructor(props) {
    super(props);
    const {conditionRedux, sortRedux, currentPage, perPageRedux} = props;
    this.state = {
      searchCondition: conditionRedux,
      page: currentPage,
      perPage: perPageRedux || NUMBER_OF_ITEMS[0].key,
      sortOption: sortRedux || SORT_OPTIONS.recommended.key,
    };
    this.firstElement = React.createRef();
    this.sortOptionList = ListHelper.generateListFromObject(SORT_OPTIONS);
  }

  componentDidMount() {
    const {list, location, history} = this.props;
    const page = CommonHelper.getValueFromQuery(location, 'page') || 1;
    this.setState({page});
    history.push(`/?page=${page}`);
    if (list.length === 0)
      this.fetchData(page);
  }

  componentDidUpdate(prevProps, prevState, snapshot) {
    const {location, currentPage} = this.props;
    const page = CommonHelper.getValueFromQuery(location, 'page') || currentPage;
    if (page !== prevState.page) {
      this.setState({page});
      this.fetchData(page);
    }
  }

  handleOptionChange = value => {
    this.setState({sortOption: value}, () => this.fetchData());
  };

  handlePerPageChange = value => {
    this.setState({perPage: value}, () => this.fetchData());
  };

  fetchData = (currentPage = null) => {
    scrollToComponent(this.firstElement.current);
    const {page, perPage, sortOption, searchCondition} = this.state;
    const {fetchBuildings} = this.props;
    fetchBuildings(currentPage || page, perPage, sortOption, searchCondition);
  };

  searchWithCondition = condition => {
    this.setState({searchCondition: condition}, () => this.fetchData());
  };

  handlePaginationChange = value => {
    const {history} = this.props;
    history.push(`/?page=${value}`);
    this.setState({page: value}, () => this.fetchData());
  };

  render() {
    const {history, loading, list, totalCount} = this.props;
    const {page, perPage, sortOption, searchCondition} = this.state;
    return (
      <Layout history={history}>
        <div className="homepage">
          <div ref={this.firstElement} />
          <Row>
            <Col span={16}>
              <Row>
                <Col span={8}>
                  {i18n.t('homepage.sort')}
                  <Select value={sortOption} className="sort-filter" defaultValue={SORT_OPTIONS.recommended.key}
                          onChange={this.handleOptionChange}
                  >
                    {this.sortOptionList.map(item => {
                      return <Select.Option key={item.key} value={item.key}>{item.value}</Select.Option>;
                    })}
                  </Select>
                </Col>
                <Col span={8}>
                  {i18n.t('homepage.number_of_items')}
                  <Select value={perPage} className="sort-filter" defaultValue={NUMBER_OF_ITEMS[0].key}
                          onChange={this.handlePerPageChange}>
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
              <SearchDetail searchWithCondition={this.searchWithCondition} initialCondition={searchCondition} />
            </Col>
          </Row>
          <Row>
            <Col className="building-list-pagination" span={16}>
              <Pagination size="small" current={page} pageSize={perPage}
                          onChange={this.handlePaginationChange}
                          total={totalCount} />
            </Col>
          </Row>
        </div>
      </Layout>
    );
  }
}


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
  location: PropTypes.object,
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
  location: PropTypes.object,
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
  location: {},
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
