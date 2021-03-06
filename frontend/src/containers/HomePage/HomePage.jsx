import React from 'react';
import PropTypes from 'prop-types';
import {Row, Col, Select, Pagination} from 'antd';
import scrollToComponent from 'react-scroll-to-component';
import {connect} from 'react-redux';

import './HomePage.scss';
import * as actions from '../../store/actions';
import {createSortOptions, createNumberOfItemOptions} from '../../constants/common';
import ListHelper from '../../helpers/list_helper';
import i18n from '../../config/i18n';
import Layout from '../../components/Layout';
import BuildingList from '../../components/BuildingList';
import ListPlaceholder from '../../components/ListPlaceholder';
import SearchDetail from './SearchDetail';
import CommonHelper from '../../helpers/common';
import FavoriteService from '../../services/favoriteService';

let SORT_OPTIONS = createSortOptions();
let NUMBER_OF_ITEMS = createNumberOfItemOptions();

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
    i18n.on('languageChanged', () => {
      SORT_OPTIONS = createSortOptions();
      NUMBER_OF_ITEMS = createNumberOfItemOptions();
      this.sortOptionList = ListHelper.generateListFromObject(SORT_OPTIONS);
    });
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

  handleFavoriteAction = (roomId, action) => {
    const {tokenData, addUserFavorite, removeUserFavorite} = this.props;
    return FavoriteService.handleFavorite(roomId, tokenData, action)
      .then(() => {
        if (action === 'create')
          addUserFavorite(roomId);
        else if (action === 'delete')
          removeUserFavorite(roomId);
      });
  };

  render() {
    const {history, loading, list, totalCount, userData, isAuthenticated} = this.props;
    const {page, perPage, sortOption, searchCondition} = this.state;
    return (
      <Layout history={history}>
        <div className="homepage">
          <div ref={this.firstElement} />
          <Row>
            <Col span={16}>
              <div className="list_option">
                <div className="list_option__sort">
                  {i18n.t('homepage.sort')}
                  <Select value={sortOption} className="sort-filter" defaultValue={SORT_OPTIONS.recommended.key}
                          onChange={this.handleOptionChange}
                  >
                    {this.sortOptionList.map(item => {
                      return <Select.Option key={item.key} value={item.key}>{item.value}</Select.Option>;
                    })}
                  </Select>
                </div>
                <div className="list_option__limit">
                  {i18n.t('homepage.number_of_items')}
                  <Select value={perPage} className="sort-filter" defaultValue={NUMBER_OF_ITEMS[0].key}
                          onChange={this.handlePerPageChange}>
                    {NUMBER_OF_ITEMS.map(item => {
                      return <Select.Option key={item.key} value={item.key}>{item.value}</Select.Option>;
                    })}
                  </Select>
                </div>
              </div>
              <Row />
              <Row>
                <div className="list">
                  {loading ? <ListPlaceholder itemCount={perPage} /> :
                    <BuildingList
                      handleFavoriteAction={this.handleFavoriteAction}
                      history={history} buildingList={list} userData={userData}
                      isAuthenticated={isAuthenticated} />
                  }
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
  userData: PropTypes.object,
  tokenData: PropTypes.object,
  isAuthenticated: PropTypes.bool,
  addUserFavorite: PropTypes.func,
  removeUserFavorite: PropTypes.func,
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
  userData: {},
  tokenData: {},
  isAuthenticated: false,
  addUserFavorite: () => {
  },
  removeUserFavorite: () => {
  },
};

const mapStateToProps = state => ({
  list: state.building.list,
  loading: state.building.loading,
  totalCount: state.building.totalCount,
  conditionRedux: state.building.condition,
  sortRedux: state.building.sort,
  currentPage: state.building.currentPage,
  perPageRedux: state.building.perPage,
  userData: state.auth.userData,
  tokenData: state.auth.tokenData,
  isAuthenticated: state.auth.isAuthenticated,
});

const mapDispatchToProps = dispatch => ({
  fetchBuildings: (page, perPage, sortOption, condition = null) => dispatch(actions.fetchBuildings(page, perPage, sortOption, condition)),
  addUserFavorite: roomId => dispatch(actions.addUserFavorite(roomId)),
  removeUserFavorite: roomId => dispatch(actions.removeUserFavorite(roomId)),
});

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
