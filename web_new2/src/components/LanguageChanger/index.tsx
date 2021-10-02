import React, { useState } from 'react';
import useTranslation from 'next-translate/useTranslation';
import setLanguage from 'next-translate/setLanguage';
import ViIcon from 'src/assets/svg/vietnam_round.svg';
import EnIcon from 'src/assets/svg/uk_round.svg';
import JaIcon from 'src/assets/svg/japan_round.svg';
import ClickAwayListener from 'src/components/ClickAwayListener';
import Popper from 'src/components/Popper';

type SupportLanguage = 'ja' | 'en' | 'vi';

interface Props {
  currentLanguage: SupportLanguage;
}

const LanguageChanger: React.FC<Props> = ({ currentLanguage }) => {
  const [anchor, setAnchor] = useState<HTMLElement | null>(null);

  const { t, lang } = useTranslation('language');
  let displayElement = null;
  const jaElement = (
    <div className="flex flex-row items-center">
      <span className="w-6">
        <JaIcon />
      </span>
      <span className="ml-2 text-sm">{t('japanese')}</span>
    </div>
  );

  const enElement = (
    <div className="flex flex-row items-center">
      <span className="w-6">
        <EnIcon />
      </span>
      <span className="ml-2 text-sm">{t('english')}</span>
    </div>
  );

  const viElement = (
    <div className="flex flex-row items-center">
      <span className="w-6">
        <ViIcon />
      </span>
      <span className="ml-2 text-sm">{t('vietnamese')}</span>
    </div>
  );

  switch (lang) {
    case 'vi':
      displayElement = viElement;
      break;
    case 'en':
      displayElement = enElement;
      break;
    case 'ja':
      displayElement = jaElement;
      break;
  }
  return (
    <ClickAwayListener onClickAway={() => setAnchor(null)}>
      <div
        className="flex"
        onClick={(e) => {
          setAnchor(anchor === null ? e.currentTarget : null);
        }}
      >
        <div className="flex flex-row px-2 cursor-pointer">
          {displayElement}
        </div>

        <Popper
          anchor={anchor}
          modifiers={{ offset: { offset: '20, 5' } }}
          open={anchor !== null}
          placement="bottom-end"
        >
          <div className="flex flex-col space-y-2 p-2">
            <div className="cursor-pointer" onClick={() => setLanguage('en')}>
              {enElement}
            </div>
            <div className="cursor-pointer" onClick={() => setLanguage('ja')}>
              {jaElement}
            </div>
            <div className="cursor-pointer" onClick={() => setLanguage('vi')}>
              {viElement}
            </div>
          </div>
        </Popper>
      </div>
    </ClickAwayListener>
  );
};

export default LanguageChanger;
