import loadable from "loadable-components";

import LoadingIndicator from '../../components/LoadingIndicator';

export default loadable(() => import('./index.js'), {
    LoadingComponent: LoadingIndicator
});
