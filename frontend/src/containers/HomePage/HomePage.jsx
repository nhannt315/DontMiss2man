import React, {PureComponent} from 'react';
import PropTypes from 'prop-types';
import {Row, Col, Select} from 'antd';

import './HomePage.scss';
import {SORT_OPTIONS, NUMBER_OF_ITEMS} from '../../constants/common';
import ListHelper from '../../helpers/list_helper';
import i18n from '../../config/i18n';
import BuildingList from '../../components/BuildingList';

const DUMMY_DATA = [
  {
    id: 1,
    name: 'M.A TOWER ROPPONGI',
    address: '東京都港区六本木３',
    access: ['東京メトロ日比谷線/六本木駅 歩8分', '東京メトロ南北線/六本木一丁目駅 歩5分', '東京メトロ南北線/麻布十番駅 歩14分'],
    photo_url: 'https://img01.suumo.com/front/gazo/fr/bukken/319/100176678319/100176678319_gw.jpg',
    structure_type: '鉄骨',
    year_built: '06/2019',
    rooms: [
      {
        id: 1,
        suumo_id: '100176658093',
        rent_fee: 152000,
        shikikin: 152000,
        reikin: 152000,
        management_cost: 10000,
        caution_fee: 0,
        size: 28.96,
        direction: '南',
        layout: '1LDK',
        floor: 3,
        facilities: 'バストイレ別、エアコン、クロゼット、フローリング、TVインターホン、浴室乾燥機、オートロック、室内洗濯置、陽当り良好、シューズボックス、システムキッチン、南向き、追焚機能浴室、角住戸、温水洗浄便座、脱衣所、エレベーター、洗面所独立、洗面化粧台、2口コンロ、宅配ボックス、外壁タイル張り、即入居可、2面採光、防犯カメラ、IHクッキングヒーター、分譲賃貸、オートバス、全居室洋室、保証人不要、二人入居相談、24時間緊急通報システム、デザイナーズ、2沿線利用可、ディンプルキー、眺望良好、築2年以内、24時間換気システム、南面リビング、耐震構造、未入居、3駅以上利用可、駅徒歩10分以内、24時間ゴミ出し可、敷地内ごみ置き場、都市ガス、洗面所にドア、南面バルコニー、高速ネット対応、保証会社利用可、初期費用カード決済可、通風良好',
        car_park: 'none',
        condition: '二人入居可',
        deal_type: '仲介',
        move_in_time: '-',
        layout_detail: null,
        damage_insurance: '2万円2年',
        guarantor: '保証会社利用必 初回月額総賃料の40％～',
        other_fees: '安心サポート料月額1000円',
        note: '巡回管理',
        updated_at: '2019/11/19',
        images: [
          {
            tag: '間取り図',
            url: 'https://img01.suumo.com/front/gazo/fr/bukken/093/100176658093/100176658093_co.jpg',
          },
        ],
      },
    ],
  },
];

class HomePage extends PureComponent {
  render() {
    const sortOptionList = ListHelper.generateListFromObject(SORT_OPTIONS);
    return (
      <div className="homepage">
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
                <Select className="sort-filter" defaultValue={NUMBER_OF_ITEMS[0].key}>
                  {NUMBER_OF_ITEMS.map(item => {
                    return <Select.Option key={item.key} value={item.key}>{item.value}</Select.Option>;
                  })}
                </Select>
              </Col>
            </Row>
            <Row />
            <Row>
              <div className="list">
                <BuildingList buildingList={DUMMY_DATA} />
              </div>
            </Row>
          </Col>
          <Col span={8} />
        </Row>
      </div>
    );
  }
}

HomePage.propTypes = {};

export default HomePage;
