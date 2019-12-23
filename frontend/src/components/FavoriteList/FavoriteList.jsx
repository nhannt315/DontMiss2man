import React, {useState, useEffect} from 'react';
import PropTypes from 'prop-types';
import i18n from '../../config/i18n';
import './FavoriteList.scss';
import FavoriteService from '../../services/favoriteService';
import OverlayLoader from '../OverlayLoader';
import FavoriteItem from './FavoriteItem';

const FavoriteList = ({tokenData, userFavoriteIds, removeUserFavorite, show}) => {
  const [list, setList] = useState([]);
  const [loading, setLoading] = useState(false);
  useEffect(() => {
    if (show) {
      setLoading(true);
      FavoriteService.getFavoriteList(tokenData)
        .then(response => {
          setList(response.data);
          setLoading(false);
        })
        .catch(() => setLoading(false));
    }
  }, [tokenData, userFavoriteIds, show]);

  const deleteFavorite = roomId => {
    return FavoriteService.handleFavorite(roomId, tokenData, 'delete')
      .then(() => {
        const newList = list.filter(e => e.id !== roomId);
        setList(newList);
        removeUserFavorite(roomId);
      });
  };
  return (
    <div className="favorite_list">
      {loading && <OverlayLoader />}
      {list.map(item => (
        <FavoriteItem key={item.id} room={item} deleteFavorite={deleteFavorite} />
      ))}
      {list.length === 0 && (
        <div className="favorite_list__no_item">
          <span>{i18n.t('common.no_item')}</span>
        </div>
      )}
    </div>
  );
};

FavoriteList.propTypes = {
  tokenData: PropTypes.object,
  userFavoriteIds: PropTypes.array,
  removeUserFavorite: PropTypes.func,
  show: PropTypes.bool,
};

FavoriteList.defaultProps = {
  show: false,
  tokenData: {},
  userFavoriteIds: [],
  removeUserFavorite: () => {
  },
};

export default FavoriteList;
