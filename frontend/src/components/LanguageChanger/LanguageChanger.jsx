import React from 'react';
import PropTypes from 'prop-types';
import {Dropdown, Menu} from 'antd';
import i18n from '../../config/i18n';
import ViIcon from '../../assets/svg/vietnam_round.svg';
import EnIcon from '../../assets/svg/uk_round.svg';
import JaIcon from '../../assets/svg/japan_round.svg';
import './LanguageChanger.scss';


const LanguageChanger = ({currentLng}) => {
  let displayElement = null;
  const jaElement = <div>
    <img className="language_changer__icon" src={JaIcon} alt="ja_icon" />
    <span className="language_changer__text">{i18n.t('language.japanese')}</span>
  </div>;

  const enElement = <div>
    <img className="language_changer__icon" src={EnIcon} alt="en_icon" />
    <span className="language_changer__text">{i18n.t('language.english')}</span>
  </div>;

  const viElement = <div>
    <img className="language_changer__icon" src={ViIcon} alt="vi_icon" />
    <span className="language_changer__text">{i18n.t('language.vietnamese')}</span>
  </div>;
  switch (currentLng) {
    default:
    case 'ja':
      displayElement = jaElement;
      break;
    case 'en':
      displayElement = enElement;
      break;
    case 'vi':
      displayElement = viElement;
      break;
  }
  const menu = (
    <Menu onClick={({key}) => i18n.changeLanguage(key)}>
      <Menu.Item key="ja">{jaElement}</Menu.Item>
      <Menu.Item key="en">{enElement}</Menu.Item>
      <Menu.Item key="vi">{viElement}</Menu.Item>
    </Menu>
  );
  return (
    <Dropdown className="language_changer" overlay={menu} placement="bottomRight">
      {displayElement}
    </Dropdown>
  );
};


LanguageChanger.propTypes = {
  currentLng: PropTypes.string,
};

LanguageChanger.defaultProps = {
  currentLng: 'ja',
};


export default LanguageChanger;
