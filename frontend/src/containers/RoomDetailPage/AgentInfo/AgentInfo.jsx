import React from 'react';
import PropTypes from 'prop-types';
import {Col, Row} from 'antd';
import Title from '../../../components/Title';
import i18n from '../../../config/i18n';

const AgentInfo = ({room}) => {
  return (
    <div>
      <Title content={i18n.t('roomDetail.current_agent')} />
      <div className="agent-detail">
        <h3>{room.agent.name}</h3>
        <Row className="agent-detail-inner">
          <Col span={4}>
            <img src={room.agent.photo_url} alt={room.agent.name} />
          </Col>
          <Col span={20}>
            <div className="agent-slogan">{room.agent.slogan}</div>
            <Row className="agent-info">
              <Col span={6} className="info-col first-col">{room.agent.address}</Col>
              <Col span={6} className="info-col">{room.agent.working_time}</Col>
              <Col span={6} className="info-col">{room.agent.access}</Col>
              <Col span={6} className="info-col">{room.agent.telephone_number}</Col>
            </Row>
          </Col>
        </Row>
      </div>
    </div>
  );
};

AgentInfo.propTypes = {
  room: PropTypes.object,
};

AgentInfo.defaultProps = {
  room: {},
};

export default AgentInfo;
