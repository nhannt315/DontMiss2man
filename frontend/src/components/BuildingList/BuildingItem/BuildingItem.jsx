import React from 'react';
import PropTypes from 'prop-types';
import {Col, Row} from 'antd';

import './BuildingItem.scss';
import RoomList from '../../RoomList';

const BuildingItem = ({item, history}) => {
  return (
    <div className="buildingitem">
      <Row className="buildingitem-detail">
        <Col className="buildingitem_object" span={5}>
          <img alt={item.name} src={item.photo_url} />
        </Col>
        <Col className="buildingitem_content" span={18} offset={1}>
          <div className="buildingitem_content-label">
            <span>賃貸マンション</span>
          </div>
          <div className="buildingitem_content-title">
            {item.name}
          </div>
          <div className="buildingitem_content-body">
            <Row className="buildingitem_detail">
              <Col className="buildingitem_detail-col1" span={6}>
                {item.address}
              </Col>
              <Col className="buildingitem_detail-col2" span={14}>
                {item.access.map(ele => <div key={ele}>{ele}</div>)}
              </Col>
              <Col className="buildingitem_detail-col3" span={4}>
                {item.year_built}
              </Col>
            </Row>
          </div>
        </Col>
      </Row>
      <Row className="buildingitem-rooms">
        <RoomList list={item.rooms} history={history} />
      </Row>
    </div>
  );
};

BuildingItem.propTypes = {
  item: PropTypes.object,
  history: PropTypes.object.isRequired,
};

BuildingItem.defaultProps = {
  item: {},
};

export default BuildingItem;
