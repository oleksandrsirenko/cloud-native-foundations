Quick Start
===========


.. _setup_env:

Install Tools
-------------

1. `VS Code <https://code.visualstudio.com>`__ (or another IDE of your choice)
2. `Python <https://www.python.org/downloads/>`__
3. `Git <https://git-scm.com/downloads>`__


Obtain Data
-----------

1. On GitHub, navigate to the `workshop repository <https://github.com/oleksandrsirenko/cloud-native-foundations>`__
2. Fork the workshop repository.
3. On GitHub, navigate to **your fork** of the workshop repository and copy the URL.
4. Clone forked repository to your local machine using git clone command. 
   It will look like this, with your GitHub username instead of ``YOUR_USERNAME``:

.. code-block:: shell

   git clone https://github.com/YOUR_USERNAME/cloud-native-foundations

See details on `how to fork and clone the repository. <https://docs.github.com/en/get-started/quickstart/fork-a-repo>`__


.. _gen_docs:


Generate Documentation
----------------------

The documentation is the core of the workshop. You can generate it automatically 
with a single line of code and use it locally without internet access. Moreover, 
you can edit and improve documentation, contributing to this project, or build 
your own knowledge base. To make it possible just follow the instructions:

1. Open the workshop folder with IDE and run terminal
2. Create the virtual environment: ``python3 -m venv venv``
3. Activate the virtual environment: ``source venv/bin/activate``
4. Install dependencies: ``pip install -r requirements.txt``
5. Change the working directory to docs: ``cd docs``
6. Generate a local copy of workshop documentation by running ``make html``
7. Open the local copy of the documentation in a web browser: ``firefox build/html/index.html``

Now you have everything you need to get the most out of this workshop.


Download Documentation
----------------------

You can download the Cloud Native Foundations Workshop docs in the following formats:

- `PDF`_ 
- `ePub`_
- `Zipped HTML`_ 
    
.. _PDF: https://cloud-native-foundations-workshop.readthedocs.io/_/downloads/en/latest/pdf/
.. _ePub: https://cloud-native-foundations-workshop.readthedocs.io/_/downloads/en/latest/epub/
.. _Zipped HTML: https://cloud-native-foundations-workshop.readthedocs.io/_/downloads/en/latest/htmlzip/


.. _how_to_use:

How To Use Workshop
-------------------

At the moment, the workshop includes two main sections: exercises and solutions. 
Follow the instructions in each exercise and try to solve the problems.

Explore the :doc:`resources` to brush up on the current topic, and feel free to 
check out the solutions.