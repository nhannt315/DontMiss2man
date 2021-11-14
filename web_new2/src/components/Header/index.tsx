import React from 'react';
import Link from 'next/link';
import { useRouter } from 'next/router';
import useTranslation from 'next-translate/useTranslation';
import LanguageChanger from 'src/components/LanguageChanger';
import { useAuth } from 'src/hooks/auth';

interface Props {
  show: boolean;
}

const Header: React.FC<Props> = ({ show }) => {
  const { t } = useTranslation('auth');
  const router = useRouter();
  const { email, token } = useAuth();
  return (
    <div
      className="flex flex-row w-full bg-white border-1 shadow-sm py-2"
      style={{ display: show ? '' : 'none' }}
    >
      <div className="container mx-auto flex flex-row">
        <div className="flex flex-row w-full items-center">
          <Link href="/">
            <img src="/logo.png" alt="Logo" className="w-10" />
          </Link>
          <span className="text-2xl text-blue-500 ml-2">DM2M</span>
          <span className="text-xs ml-2">{t('common:slogan')}</span>
          <div className="ml-auto flex flex-row">
            <LanguageChanger currentLanguage={'ja'} />
            {!(email && token) && (
              <button
                className="border-2 border-solid text-sm border-gray-200 text-gray-500 bg-white px-2 py-1 rounded-md hover:border-blue-300 hover:text-blue-500"
                onClick={() => router.push('/login')}
              >
                {t('login')}
              </button>
            )}
            {email && token && <div>{email}</div>}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Header;