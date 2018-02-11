===============================
Get Instagram User Media in Go_
===============================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/goigmedia?status.png
   :target: https://godoc.org/github.com/siongui/goigmedia

.. image:: https://api.travis-ci.org/siongui/goigmedia.png?branch=master
   :target: https://travis-ci.org/siongui/goigmedia

.. image:: https://goreportcard.com/badge/github.com/siongui/goigmedia
   :target: https://goreportcard.com/report/github.com/siongui/goigmedia

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://raw.githubusercontent.com/siongui/goigmedia/master/UNLICENSE

.. image:: https://img.shields.io/badge/Status-Beta-brightgreen.svg

.. image:: https://img.shields.io/twitter/url/https/github.com/siongui/goigmedia.svg?style=social
   :target: https://twitter.com/intent/tweet?text=Wow:&url=%5Bobject%20Object%5D


Get links of Instagram_ user media (photos and videos) in Go.


Obtain Cookies
++++++++++++++

The following three values are must to access the Instagram API.

- ``ds_user_id``
- ``sessionid``
- ``csrftoken``

First login to Instagram_ from Chrome browser, and there are two ways to get the
above information:

1. From `Chrome Developer Tools`_: See this `SO answer`_ or `Obtain cookies`_
   section in `instastories-backup`_ repo.

.. image:: https://i.stack.imgur.com/psJLZ.png
   :align: center
   :alt: ds_user_id sessionid csrftoken

2. From Chrome extension: Use EditThisCookie_ or `cookie-txt-export`_ or other
   cookie tools.


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `GitHub - siongui/goiguserid: Get id of Instagram user in Go <https://github.com/siongui/goiguserid>`_
.. [2] `GitHub - siongui/goigstorylink: Get Links (URL) of Instagram Stories in Go <https://github.com/siongui/goigstorylink>`_
.. [3] `GitHub - siongui/goigfollow: Get Instagram following and followers in Go <https://github.com/siongui/goigfollow>`_
.. [4] `GitHub - siongui/goigstorydl: Download Instagram Stories in Go <https://github.com/siongui/goigstorydl>`_


.. _Go: https://golang.org/
.. _Instagram: https://www.instagram.com/
.. _Chrome Developer Tools: https://developer.chrome.com/devtools
.. _SO answer: https://stackoverflow.com/a/44773079
.. _Obtain cookies: https://github.com/hoschiCZ/instastories-backup#obtain-cookies
.. _instastories-backup: https://github.com/hoschiCZ/instastories-backup
.. _EditThisCookie: https://www.google.com/search?q=EditThisCookie
.. _cookie-txt-export: https://github.com/siongui/cookie-txt-export.go
.. _UNLICENSE: http://unlicense.org/
